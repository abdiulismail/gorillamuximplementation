package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mux/models"
	"net/http"
)

func DeleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")

	params := mux.Vars(r)

	for index, item := range models.MyBooks{
		if item.ID == params["id"]{
			models.MyBooks = append(models.MyBooks[:index],models.MyBooks[index+1:]...)
			break
		}
		_ = json.NewEncoder(w).Encode(models.MyBooks)
	}

}
