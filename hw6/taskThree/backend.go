package main

import (
	"fmt"
	"log"
	"net/http"
)

func form(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, "form.html")
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		address := r.FormValue("address")
		token := name + ":" + address
		cookie := http.Cookie{
			Name:  "token",
			Value: token,
		}
		http.SetCookie(w, &cookie)
		http.ServeFile(w, r, "form.html")
	default:
		//gfmt.Fprintf(w, "method not supported")
		err := fmt.Errorf("Error. Method not supported")
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/", form)

	fmt.Printf("Starting server for testing HTTP POST\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
