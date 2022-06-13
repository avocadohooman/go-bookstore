package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github/avocadohooman/go-bookstore/pkg/utils"
	"github/avocadohooman/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId)
	if err != nil {
		fmt.Println("Invalid ID", ID)
		w.WriteHeader(http.StatusBadRequest)
	}
	findBook, _ := models.GetBookById(ID)
	res, err := json.Marshal(findBook)
	if err != nil {
		fmt.Println("Converting to json failed")
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}