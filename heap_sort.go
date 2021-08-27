package main

import (
	"fmt"
)

func main(){
	data := []int{7777,1000,  51, 200, -12, 9929, 1450, 1680,  44, 211}

	fmt.Printf("Original data - %d\n",data)

	result := HeapSort(data)

	fmt.Printf("Sorted data - %d\n",result)


}


func HeapSort(data []int) []int {
	heapify(data)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		siftDown(data, 0, i)
	}
	return data
}

func heapify(data []int) {
	for i := (len(data) - 1) / 2; i >= 0; i-- {
		siftDown(data, i, len(data))
	}
}

func siftDown(heap []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}

	}
}