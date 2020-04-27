package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthday  time.Time
}

type People []Person

func (p People) Len() int {
	return len(p)
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p People) Less(i, j int) bool {
	if p[i].birthday.Unix() == p[j].birthday.Unix() {
		if p[i].firstName == p[j].firstName {
			return p[i].lastName < p[j].lastName
		}
		return p[i].firstName < p[j].firstName
	}
	return p[i].birthday.Unix() > p[j].birthday.Unix()
}

func main() {

	groupB := People{
		{"a", "c", time.Date(2000, 0, 1, 12, 0, 0, 0, time.UTC)},
		{"a", "b", time.Date(2000, 0, 1, 12, 0, 0, 0, time.UTC)},
		{"a", "a", time.Now()},
	}
	sort.Sort(groupB)
	fmt.Println(groupB)
}
