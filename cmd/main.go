package main

import (
	"github.com/gorilla/mux"
	"log"
	"mux/api"
	"mux/models"
	"net/http"
)

func main(){
	//init router
	r := mux.NewRouter()
	
	//mock data
	models.MyBooks = append(models.MyBooks,models.Book{
		ID:"1",
		Isbn:   "344",
		Title:  "book1",
		Author: &models.Author{
			FirstName: "john",
			LastName:  "doe",
		},})

	models.MyBooks = append(models.MyBooks,models.Book{
		ID:"2",
		Isbn:   "84848",
		Title:  "book2",
		Author: &models.Author{
			FirstName: "allan",
			LastName:  "doe",
		},})


	//route handlers/ endpoints
	r.HandleFunc("/api/books",api.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",api.GetBook).Methods("GET")
	r.HandleFunc("/api/books",api.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",api.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}",api.DeleteBooks).Methods("DELETE")

	log.Fatal(	http.ListenAndServe(":8000",r))
}