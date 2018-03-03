package main

//Created by Nick Regan
//With assistance from Matt Silverlock's tutoiral:
//	http://blog.questionable.services/article/testing-http-handlers-go/

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

//Basic test function. Verifies basic functionality of the handler.
//Given more time, I would write tests for each module.
func TestRedditSearch(t *testing.T) {
    request, err := http.NewRequest("GET", "/redditer/_", nil)
    if err != nil {
        t.Fatal(err)
    }

    responseRecorder := httptest.NewRecorder()
    handler := http.HandlerFunc(RedditSearch)

    handler.ServeHTTP(responseRecorder, request)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }
}