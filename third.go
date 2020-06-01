package main

import (
	"fmt"
	"sync"
)

var sl []string
var wg sync.WaitGroup
var mutex = &sync.Mutex{}

func addLine(words string) {
	mutex.Lock()
	sl = append(sl, words)
	mutex.Unlock()
	wg.Done()
}

func main() {
	wg.Add(4)

	go addLine("I'll")
	go addLine(" be here")
	go addLine(" all day")
	go addLine(" and you'll be too")
	wg.Wait()
	fmt.Println(sl)
}
