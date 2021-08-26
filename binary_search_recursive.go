package main

import (
	"fmt"
)

func main(){
	data := []int{-12, 44, 51, 200, 211, 1000, 1450, 1680, 7777, 9929}
	numberToSearchFor := 1450

	fmt.Printf("Original sorted data - %d\n",data)
	result,searchCount := binarySearchRecursion(data, numberToSearchFor)
	if result == -1 {
		fmt.Printf("Item %d not found\n",numberToSearchFor)
		return
	}

	fmt.Printf("Item %d found at index %d after %d searche(s)\n",numberToSearchFor, result, searchCount)

}

func binarySearchRecursion(a []int, search int) (result int, searchCount int) {
	fmt.Println("Binary search")
	fmt.Printf("Original item: %d\n", a)
	fmt.Printf("Item to look for: %d\n", search)
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = binarySearchRecursion(a[:mid], search)
		fmt.Println("recursive call")
	case a[mid] < search:
		result, searchCount = binarySearchRecursion(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
		fmt.Println("recursive call")
	default: // a[mid] == search
		result = mid // found
		fmt.Println("found")
	}
	searchCount++
	return
}