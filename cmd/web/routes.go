package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/full/:shortURL", http.HandlerFunc(app.getFullURL))
	mux.Get("/shorten/:url", http.HandlerFunc(app.shortenURL))

	return standardMiddleware.Then(mux)
}
