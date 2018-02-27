package main

import (
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"fmt"
	
	"github.com/gorilla/mux"
)

func RedditSearch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    search := vars["search"]
    
	url := "http://reddit.com/r/ProgrammerHumor/top/.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "TDI-project")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	page := Page{}
	jsonErr := json.Unmarshal(body, &page)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	
	fmt.Fprintln(w, "Search:", search)
	fmt.Fprintln(w, page.Data.Children[0].Data)
}