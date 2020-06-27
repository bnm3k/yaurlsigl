package main

import (
	"encoding/json"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	w.Write([]byte("Yet another URL shortener"))
}

func (app *application) getFullURL(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":shortURL")
	originalURL, err := app.store.GetFullURL(id)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "")
	http.Redirect(w, r, "http://"+originalURL, http.StatusFound)
}

// ShortenResult ...
type ShortenResult struct {
	OriginalURL string `json:"original"`
	Shortcode   string `json:"shortcode"`
}

func (app *application) shortenURL(w http.ResponseWriter, r *http.Request) {
	fullURL := r.URL.Query().Get(":url")
	shortcode, err := app.store.ShortenURL(fullURL)
	if err != nil {
		// TODO: add check for whether it's server
		// or client error
		app.clientError(w, http.StatusBadRequest)
		return
	}
	res := &ShortenResult{OriginalURL: fullURL, Shortcode: shortcode}
	resJSON, err := json.Marshal(res)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJSON)
}
