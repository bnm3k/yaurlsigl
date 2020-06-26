package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	w.Write([]byte("Hello there"))
}

func (app *application) getFullURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	w.Write([]byte(fmt.Sprintf("full URL: %s\n", id)))
}
