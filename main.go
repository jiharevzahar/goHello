package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/timetrackerdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/groups", GetGroups).Methods("GET")
	router.HandleFunc("/groups/", CreateGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", UpdateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", DeleteGroup).Methods("DELETE")

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	router.HandleFunc("/timeframes", createTimeframe).Methods("POST")
	router.HandleFunc("/timeframes/{id}", deleteTimeframe).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

