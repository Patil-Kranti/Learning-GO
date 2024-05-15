package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Patil-Kranti/Learning-GO/tree/main/bookstore-management-using-postgress/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allbooks := models.GetAllBooks()
	json.NewEncoder(w).Encode(allbooks)
}
