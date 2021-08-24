package main

import (
	"fmt"
)

func main(){
	data := []int{9, 14, 3, 1, 0, 44, 100, 12}

	fmt.Printf("Original data - %d\n",data)
	SelectionSort(data)
	fmt.Printf("Sorted data - %d\n",data)
}

func SelectionSort(data []int) {
	length := len(data)
	for i := 0; i < length; i++ {
		maxIndex := 0
		for j := 1; j < length-i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			}
		}
		data[length-i-1], data[maxIndex] = data[maxIndex], data[length-i-1]
	}
}