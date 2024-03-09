package main

import "net/http"


func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", app.homeHandle)
	mux.HandleFunc("GET /snippet", app.showSnippetHandle)
	mux.HandleFunc("POST /snippet/create", app.createSnippetHandle)

	fileServer := http.FileServer(neueredFileSystem{http.Dir("./interface/static/")})

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	return mux
}