package controllers

import (
	"encoding/json"
	"fmt"
	"github/avocadohooman/go-bookstore/pkg/models"
	"github/avocadohooman/go-bookstore/pkg/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
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

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Invalid ID", ID)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid id")
	}
	book := models.DeleteBook(ID)
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(book)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)
	params := mux.Vars(r)
	id := params["id"]
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Invalid ID", ID)
		w.WriteHeader(http.StatusBadRequest)
		res := make(map[string]string)
		res["error"] = "Invalid id:" + id
		jsonRes, err := json.Marshal(res)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonRes)
	}
	b := UpdateBook.UpdateBook(ID)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}