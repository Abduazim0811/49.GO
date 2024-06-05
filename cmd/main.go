package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Greeting struct {
	Message string `json:"message"`
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	name := queryValues.Get("name")
	if name == "" {
		name = "world"
	}
	client:=http.Client{
		Timeout: 5*time.Second,
	}
	stringurl:="http://localhost:9001/greeting/"+name
	resp,err:=client.Get(stringurl)
	if err!=nil{
		log.Fatal(err)
	}

	if resp.StatusCode!=http.StatusOK{
		log.Fatal(err)
	}
	m:=map[string]any{}
	err=json.NewDecoder(resp.Body).Decode(&m)
	if err!=nil{
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m)
}

func main() {
	port := ":7777"
	http.HandleFunc("/greet", greetHandler)

	http.ListenAndServe(port, nil)
}
