package main

import "fmt"

func main(){
	dataInput := []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
	fmt.Printf("Unsorted items: %d\n", dataInput)
	sortedItems := MergeSort(dataInput)
	fmt.Printf("Sorted items: %d\n", sortedItems)
}

func MergeSort(unsorted []int) []int{
	//fmt.Println("Merge Sort")
	//fmt.Printf("Unsorted items: %d\n", unsorted)
	if len(unsorted) < 2 {
		return unsorted
	}
	mid := (len(unsorted)) / 2
	return Merge(MergeSort(unsorted[:mid]), MergeSort(unsorted[mid:]))
}

func Merge(left, right []int) []int{

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}