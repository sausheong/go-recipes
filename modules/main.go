package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sausheong/go-recipes/modules/app"
)

var dir string

func main() {
	server()
}

func server() {
	r := mux.NewRouter()
	r.HandleFunc("/{category}", app.Index)

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
	}
	log.Println("Starting server at", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
