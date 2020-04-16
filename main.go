package main

import "fmt"

func median(arr [5]int) float64 {
	n := len(arr)
	var i, j, a int

	for i = 0; i < n; i++ {
		for j = i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				a = arr[i]
				arr[i] = arr[j]
				arr[j] = a
			}
			if arr[i] > arr[j] {
				a = arr[i]
				arr[i] = arr[j]
				arr[j] = a
			}
		}
	}

	if n%2 == 0 {
		return (float64(arr[n/2]) + float64(arr[n/2-1])) / 2.0
	} else {
		return float64(arr[n/2])
	}
}

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 25, 50}
	fmt.Println(median(numbers))
}
