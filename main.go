package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Get("/", home)
	r.Get("/item", listItems)
	r.Get("/item/{id:\\d}", getItem)
	r.Post("/item", addItem)
	r.Delete("/item/{id:\\d}", delItem)
	http.ListenAndServe(":5000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	http.ServeFile(w, r, "go.mod")
}

/*
type Item struct {
	Desc string `json:"desc"`
	Qty int `json:"qty"`
	Price float64 `json:"price"`
	AddedAt Time `json:"added_at"`
}
*/
