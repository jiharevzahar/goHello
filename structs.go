package main

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
