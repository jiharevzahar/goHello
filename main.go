package main

import (
	"fmt"
	"sort"
)

func median(arr []int) float64 {
	n := len(arr)

	sort.Ints(arr)

	if n%2 == 0 {
		return (float64(arr[n/2]) + float64(arr[n/2-1])) / 2.0
	} 
	return float64(arr[n/2])
	
}

func main() {
	var odd []int = []int{1, 2, 3, 25, 50}
	fmt.Println(median(odd))

	var even []int = []int{5, 17, 3, 9, 14, 2}
	fmt.Println(median(even))
}
