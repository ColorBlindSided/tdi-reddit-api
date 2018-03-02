package main

//Created by: Nick Regan
//With assistance from Cory Lanou's tutorial:
//	https://thenewstack.io/make-a-restful-json-api-go/

import (
    "log"
    "net/http"
)

//Entry point
func main() {
    
	router := NewRouter()
	
	log.Fatal(http.ListenAndServe(":8080", router))
}