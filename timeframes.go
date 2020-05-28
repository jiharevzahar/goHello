package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func createTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	task_id := keyVal["task_id"]
	from_time := keyVal["from_time"]
	to_time := keyVal["to_time"]

	stmt, err := db.Prepare("INSERT INTO timeframes(task_id, from_TIme, to_Time) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(task_id, from_time, to_time)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New timeframe was created")
}

func deleteTimeframe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM timeframes WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Timeframe with ID = %s was deleted", params["id"])
}
