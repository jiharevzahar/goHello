package main

import "fmt"

//
//In my case it's not nessesary to return float64 due to the fact that i select the median from the array.
//If it was meant to find the median(not from the array), then the implementation would be much simplier
//

func median(arr [5]int) float64 {
	size := len(arr)
	for i := 0; i < size; i++ {
		var (
			n int = 0
			m int = 0
		)
		var j = 0
		for ; j < size; j++ {
			if arr[j] != arr[i] {
				if arr[j] > arr[i] {
					n++
				} else {
					m++
				}
			}
		}
		if n == m {
			return float64(arr[i])
		}
	}
	return 0
}

func main() {
	var numbers [5]int = [5]int{1, 2, 3, -5, 5}
	fmt.Println(median(numbers))
}
