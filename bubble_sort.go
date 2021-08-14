package main

import "fmt"

func main(){
	dataInput := []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
	bubbleSort(dataInput)
}

func bubbleSort(unsorted []int){
	fmt.Println("Bubble Sort")
	fmt.Printf("Unsorted items: %d\n", unsorted)
	fmt.Println()
	var i, j int
	for i = 0; i < (len(unsorted) - 1); i++ {
		for j = 1; j < (len(unsorted) - i); j++ {
			if unsorted[j-1] > unsorted[j] {
				temp := unsorted[j]
				unsorted[j] = unsorted[j-1]
				unsorted[j-1] = temp
			}
		}
		fmt.Printf("Unsorted array after %d pass is %d\n:", i+1, unsorted)
	}
}
