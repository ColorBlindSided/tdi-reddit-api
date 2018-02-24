package main

import (
    "fmt"
    //"html"
    "log"
    "net/http"
	
	"github.com/gorilla/mux"
)

func main() {
    
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/redditer/{search}", RedditSearch)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func RedditSearch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    search := vars["search"]
    fmt.Fprintln(w, "Search:", search)
}