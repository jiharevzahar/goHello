package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type request struct {
	Host       string  `json:"host"`
	UserAgent  string  `json:"user_agent"`
	RequestURL string  `json:"request_uri"`
	Headers    headers `json:"headers"`
}

type headers struct {
	Accept    []string `json:"Accept"`
	UserAgent []string `json:"User-Agent"`
}

func main() {
	req := &request{
		Host:       "127.0.0.1:8080",
		UserAgent:  "curl/7.67.0",
		RequestURL: "/anything/you?want",
		Headers: headers{
			Accept:    []string{"*/*"},
			UserAgent: []string{"curl/7.67.0"},
		},
	}

	js, err := json.Marshal(*req)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(js)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Panic(err)
	}
}

