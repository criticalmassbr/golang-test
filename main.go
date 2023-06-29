package main

import (
	"DialogBookStore/configs"
	"DialogBookStore/handlers"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	// Books Routers
	r := chi.NewRouter()
	r.Post("/", handlers.BookCreate)
	r.Put("/{id}", handlers.BookUpdate)
	r.Delete("/{id}", handlers.BookDelete)
	r.Get("/", handlers.BookList)
	r.Get("/{id}", handlers.Get)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
