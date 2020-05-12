package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Host       string              `json:"host"`
	UserAgent  string              `json:"user_agent"`
	RequestURL string              `json:"request_uri"`
	Header     map[string][]string `json:"headers"`
}

func jsonOutput(w http.ResponseWriter, r *http.Request) {
	var outputRequest request
	outputRequest.Host = r.Host
	outputRequest.UserAgent = r.UserAgent()
	outputRequest.RequestURL = r.RequestURI
	outputRequest.Header = r.Header

	output, err := json.Marshal(outputRequest)
	if err != nil {
		log.Panic(err)
	}
	w.Write(output)
}

func main() {

	http.HandleFunc("/", jsonOutput)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}
