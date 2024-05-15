package main

import (
	"log"
	"net/http"

	"github.com/Patil-Kranti/Learning-GO/tree/main/bookstore-management-using-postgress/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
