package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/salihprnz07/go-bookstore/pkg/routes"
	_ "github.com/salihprnz07/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
