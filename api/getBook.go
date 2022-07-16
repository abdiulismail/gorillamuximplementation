package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"mux/models"
	"net/http"
)

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	params := mux.Vars(r)  //get params
	//loop through books
	for _, item := range models.MyBooks{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Book{})

}
