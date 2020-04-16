package main

import (
	"fmt"
	"sort"
)

func average(arr [6]int) float64 {
	sum := 0.0
	for _, value := range arr {
		sum = sum + float64(value)
	}
	return sum / float64(len(arr))
}

func max(arr []string) string {
	maximumInd := 0
	for index, value := range arr {
		if len(arr[maximumInd]) < len(value) {
			maximumInd = index
		}
	}
	return arr[maximumInd]
}

func reverse(arr []int64) []int64 {
	var s []int64
	for i := len(arr) - 1; i >= 0; i-- {
		s = append(s, arr[i])
	}
	return s
}

func printSorted(mp map[int]string) {
	keys := make([]int, 0, len(mp))
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, mp[k])
	}
}

func main() {
	var numbers [6]int = [6]int{1, 2, 3, 4, 5, 0}
	fmt.Println(average(numbers))
	var slice = []string{"one", "two", "three", "sdds"}
	fmt.Println(max(slice))
	var mas = []int64{1, 2, 3, 4}

	result := reverse(mas)
	for _, value := range result {
		fmt.Printf("%d ", value)
	}

	mp := map[int]string{1: "aa", 3: "cc", 2: "bb"}
	printSorted(mp)
}
