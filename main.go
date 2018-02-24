package main

import (
    //"fmt"
    //"html"
    "log"
    "net/http"
	"encoding/json"
	
	"github.com/gorilla/mux"
)

type Post struct {
	Title	string	`json:"title"`
	Date	string	`json:"date"`
	Content	string	`json:"content"`
}

type Posts []Post

func main() {
    
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/redditer/{search}", RedditSearch)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func RedditSearch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    search := vars["search"]
    
	posts := Posts{
		Post{Title: search, Content: "1"},
		Post{Title: search, Content: "2"},
	}
	
	json.NewEncoder(w).Encode(posts)
}