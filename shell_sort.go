package main

import "fmt"

func main(){
	dataInput := []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
	fmt.Printf("Unsorted array: %d\n", dataInput)
	ShellSort22(dataInput)
	fmt.Printf("Sorted array: %d\n", dataInput)
}

func ShellSort22(data []int) {
	n := len(data)
	h := 1
	for h < n/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && data[j] < data[j-h]; j -= h {
				data[j], data[j-h] = data[j-h], data[j]
			}
		}
		h /= 3
	}
}
