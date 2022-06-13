package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	"github/avocadohooman/go-bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9090", router))
}
