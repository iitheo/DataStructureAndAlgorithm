package main

import (
	"fmt"
)

func main(){
	data := []int{-12, 44, 51, 200, 211, 1000, 1450, 1680, 7777, 9929}
	numberToSearchFor := 1450

	fmt.Printf("Original sorted data - %d\n",data)
	result := binarySearchIterative(data, numberToSearchFor)
	if result == -1 {
		fmt.Printf("Item %d not found\n",numberToSearchFor)
		return
	}

	fmt.Printf("Item %d found at index %d\n",numberToSearchFor, result)

}

func binarySearchIterative(nums []int, target int) int {
	fmt.Println("Binary Search iterative")
	//count := 0
	start := 0
	end := len(nums) - 1

	for(start <= end) {
		mid := (start + end)/2
		if(target == nums[mid]) {
			//count++
			//fmt.Printf("count: %d\n", count)
			return mid
		} else if (target < nums[mid]){
			//count++
			//fmt.Printf("count: %d\n", count)
			end = mid - 1
		} else {
			//count++
			//fmt.Printf("count: %d\n", count)
			start = mid + 1
		}
	}

	return -1
}