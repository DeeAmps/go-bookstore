package controllers

import (
	"encoding/json"
	"github.com/DeeAmps/go-bookstore-api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/DeeAmps/go-bookstore-api/pkg/models"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetBooks()
	res, _ := json.Marshal(books)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{}
	utils.ParseBody(r, &book)
	results := book.CreateBook()
	if results.RowsAffected > 0 {
		w.WriteHeader(http.StatusCreated)
	}
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	book := models.GetBookByID(id)
	if book.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte{})
		return
	}
	res, _ := json.Marshal(book)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	models.DeleteBook(id)
	w.WriteHeader(http.StatusNoContent)
}
