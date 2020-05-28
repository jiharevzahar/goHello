package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []Task
	resultTask, err := db.Query("SELECT id_t, title, group_id from tasks")
	if err != nil {
		panic(err.Error())
	}
	defer resultTask.Close()
	for resultTask.Next() {
		var task Task
		var timeFrames []Timeframe
		err := resultTask.Scan(&task.ID, &task.Title, &task.Group)
		if err != nil {
			panic(err.Error())
		}

		resultTimeframe, err := db.Query("SELECT from_TIme, to_Time from timeframes WHERE task_id = " + task.ID)
		if err != nil {
			panic(err.Error())
		}
		defer resultTimeframe.Close()
		for resultTimeframe.Next() {
			var timeFrame Timeframe
			err := resultTimeframe.Scan(&timeFrame.From, &timeFrame.To)
			if err != nil {
				panic(err.Error())
			}
			timeFrames = append(timeFrames, timeFrame)
		}

		task.Time = timeFrames
		tasks = append(tasks, task)
	}
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	title := keyVal["title"]
	group_id := keyVal["group_id"]

	//DELETE THIS 3 LINE OF CODE
	fmt.Fprintf(w, title+"\t")
	fmt.Fprintf(w, group_id+"\t")

	stmt, err := db.Prepare("INSERT INTO tasks(group_id, title) VALUES(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(group_id, title)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New task was added")
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE tasks SET group_id = ?, title = ? WHERE id_t = ?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	group_id := keyVal["group_id"]
	title := keyVal["title"]
	_, err = stmt.Exec(group_id, title, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Task with ID = %s was updated", params["id"])
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id_t = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Task with ID = %s was deleted", params["id"])
}