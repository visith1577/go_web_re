package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"visithflt.net/web_t/pkg/models/mysql"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {

	// cmd line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Mysql datasource name")
	flag.Parse()

	// logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)

	if err != nil {
		errLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		infoLog: infoLog,
		errLog:  errLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	infoLog.Printf("Starting Server on %s", *addr)

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errLog.Fatal(err)
		}
	}()

	<-interrupt
	infoLog.Printf("Gracefully shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	err = srv.Shutdown(ctx)
	if err != nil {
		infoLog.Println("Error during graceful shutdown:", err)
	}

	infoLog.Println("Server shutdown complete")

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
