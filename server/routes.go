package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/{url-id}", GetHandler)
	mux.Post("/", PostHandler)

	return mux
}
