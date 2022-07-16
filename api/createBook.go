package api

import (
	"encoding/json"
	"math/rand"
	"mux/models"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	var book models.Book

	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	models.MyBooks = append(models.MyBooks, book)
	json.NewEncoder(w).Encode(book)

}
