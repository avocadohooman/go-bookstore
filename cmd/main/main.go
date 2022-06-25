package main

import (
	"fmt"
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
	fmt.Println("Server running on port 9090")
	log.Fatal(http.ListenAndServe("localhost:9090", router))
}
