package main

import "fmt"

func main(){
	dataInput := []int{4, 1, 5, 8, 9, 3, 2, 6, 7}
	fmt.Printf("Unsorted array: %d\n", dataInput)
	fmt.Printf("Sorted array: %d\n", HeapSort(dataInput))
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
