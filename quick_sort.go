package main

import "fmt"

func main(){
	dataInput := []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
	fmt.Printf("Unsorted array: %d\n", dataInput)
	fmt.Printf("Sorted array: %d\n", quickSort(dataInput))
}

func quickSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	qsort(newArr, 0, len(arr)-1)

	return newArr
}

func qsort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	qsort(arr, start, splitIndex-1)
	qsort(arr, splitIndex+1, end)
	return
}
