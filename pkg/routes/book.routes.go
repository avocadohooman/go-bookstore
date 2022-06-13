package routes

import (
	"github.com/gorilla/mux"
	"github/avocadohooman/go-bookstore/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook()).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks()).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById()).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook()).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook()).Methods("DELETE")
}