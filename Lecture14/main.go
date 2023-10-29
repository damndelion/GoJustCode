package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Mount("/debug", middleware.Profiler())
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		return
	}
}
