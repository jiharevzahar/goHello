package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type Group struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Task  []Task `json:"task"`
}

type Task struct {
	ID    string      `json:"id"`
	Title string      `json:"title"`
	Group string      `json:"group"`
	Time  []Timeframe `json:"time"`
}

type Timeframe struct {
	From string `json:"from"`
	To   string `json:"to"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/timetrackerdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	router := mux.NewRouter()

	router.HandleFunc("/groups", getGroups).Methods("GET")
	router.HandleFunc("/groups/", createGroup).Methods("POST")
	router.HandleFunc("/groups/{id}", updateGroup).Methods("PUT")
	router.HandleFunc("/groups/{id}", deleteGroup).Methods("DELETE")

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")

	router.HandleFunc("/timeframes", createTimeframe).Methods("POST")
	router.HandleFunc("/timeframes/{id}", deleteTimeframe).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}

func getGroups(w http.ResponseWriter, r *http.Request) {
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

func createGroup(w http.ResponseWriter, r *http.Request) {
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

func updateGroup(w http.ResponseWriter, r *http.Request) {
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

func deleteGroup(w http.ResponseWriter, r *http.Request) {
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
