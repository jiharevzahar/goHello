package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var groups []Group
	resultTask, err := db.Query("SELECT id_j, title from groups")
	if err != nil {
		panic(err.Error())
	}
	defer resultTask.Close()
	for resultTask.Next() {
		var group Group
		var tasks []Task
		err := resultTask.Scan(&group.ID, &group.Title)
		if err != nil {
			panic(err.Error())
		}

		resultTask, err := db.Query("SELECT id_t, title from tasks WHERE group_id = " + group.ID)
		if err != nil {
			panic(err.Error())
		}
		defer resultTask.Close()
		for resultTask.Next() {
			var task Task
			err := resultTask.Scan(&task.ID, &task.Title)
			if err != nil {
				panic(err.Error())
			}
			tasks = append(tasks, task)
		}

		group.Task = tasks
		groups = append(groups, group)
	}
	json.NewEncoder(w).Encode(groups)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]

	stmt, err := db.Prepare("INSERT INTO groups(title) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New group was added")
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE groups SET title = ? WHERE id_j = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	_, err = stmt.Exec(title, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Group with ID = %s was updated", params["id"])
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM groups WHERE id_j = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Group with ID = %s was deleted", params["id"])
}