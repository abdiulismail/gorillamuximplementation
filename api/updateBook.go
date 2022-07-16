package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mux/models"
	"net/http"
)

func UpdateBooks(w http.ResponseWriter, r *http.Request){

	w.Header().Set("content-type","application/json")

	params := mux.Vars(r)

	for index, item := range models.MyBooks{
		if item.ID == params["id"]{
			models.MyBooks = append(models.MyBooks[:index],models.MyBooks[index+1:]...)

			var book models.Book

			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			models.MyBooks = append(models.MyBooks, book)
			json.NewEncoder(w).Encode(book)
			return

		}
		_ = json.NewEncoder(w).Encode(models.MyBooks)
	}

}
