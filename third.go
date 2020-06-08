package main

import (
	"fmt"
	"sync"
)

var sl []string
var wg sync.WaitGroup

func addLine(words string) {
	sl = append(sl, words)
	wg.Done()
}

func main() {

	text := []string{"I'll", " be here", " all day", " and you'll", " and you'll be too"}
	for _, str := range text {
		wg.Add(1)
		go addLine(str)
		wg.Wait()
	}

	wg.Wait()
	fmt.Println(sl)
}
