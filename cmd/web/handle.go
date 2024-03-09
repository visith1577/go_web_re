package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) homeHandle(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./interface/html/home.page.tmpl", 
		"./interface/html/base.layout.tmpl",
		"./interface/html/footer.partial.tmpl",
	} 

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errLog.Println(err.Error())
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)

	if err != nil {
		app.errLog.Println(err.Error())
		app.serverError(w, err)
	}
}

func (app *application) showSnippetHandle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display sippet with ID %d", id)
}

func (app *application) createSnippetHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusForbidden)
		return
	}

	w.Write([]byte("create snippet"))
}
