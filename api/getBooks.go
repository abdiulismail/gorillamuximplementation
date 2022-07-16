package api

import (
	"encoding/json"
	"mux/models"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	json.NewEncoder(w).Encode(models.MyBooks)
}
