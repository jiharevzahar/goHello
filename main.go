package main

import "fmt"

func median(arr [5]int) {
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
			fmt.Println("median - ", arr[i])
		}
	}
}

func main() {

	var numbers [5]int = [5]int{1, 2, 3, -5, 5}
	fmt.Println("res")
	median(numbers)

}
