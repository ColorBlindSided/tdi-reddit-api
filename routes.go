package main

import (
    "net/http"

    "github.com/gorilla/mux"
)

//A struct to store endpoint route data
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

//There's only one endpoint, but adding more would be easy
var routes = Routes{Route{"Redditer","GET","/redditer/{search}",RedditSearch}}