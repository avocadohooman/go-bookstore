package routes

import (
	"github/avocadohooman/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(router *mux.Router) {
	// router.HandleFunc("/book/", controllers.CreateBook()).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.UpdateBook()).Methods("GET")
	// router.HandleFunc("/book/{id}", controllers.UpdateBook()).Methods("DELETE")
}