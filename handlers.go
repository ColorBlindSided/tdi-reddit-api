package main

import (
	"encoding/json"
	"net/http"
	"log"
	"io/ioutil"
	"time"
	"reflect"
	"strings"
	
	"github.com/gorilla/mux"
)

func RedditSearch(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    search := vars["search"]
    
	//Set the reddit url to pull from
	url := "http://reddit.com/r/ProgrammerHumor/top/.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Be a good internet citizen
	req.Header.Set("User-Agent", "TDI-project")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	//Unmarshal downloaded JSON into Page struct
	page := Page{}
	jsonErr := json.Unmarshal(body, &page)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	
	//Loop through all the posts to search for matches
	var results Post_Collection
	for i := 0; i < len(page.Data.Children); i++ {
		current_page := page.Data.Children[i]
		v := reflect.ValueOf(current_page.Data)
		match := false
		
		//Loop through all the fields to check for matches
		for f := 0; f < v.NumField() && !match; f++ {
			value := v.Field(f).Interface()
			
			//Type assertion to ensure we're only searching string fields
			s, ok := value.(string)
			if ok && strings.Contains(s, search) {
				match = true
			}
		}
		if match {
			results.Matches = append(results.Matches, current_page)
		}
	}
	
	//Encode the results as JSON
	json.NewEncoder(w).Encode(results)
}