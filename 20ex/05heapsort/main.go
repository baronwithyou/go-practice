package main

import "fmt"

func main() {
	arr := []int{9, 4, 2, 8, 7, 1}

	n := len(arr)

	for i := 0; i < n; i++ {
		buildHeap(arr, n-i)

		arr[0], arr[n-i-1] = arr[n-i-1], arr[0]
	}

	fmt.Println(arr)
}

func heapify(arr []int, n, idx int) {
	root := idx
	left := 2*idx + 1
	right := 2*idx + 2

	if left < n && arr[left] > arr[root] {
		root = left
	}

	if right < n && arr[right] > arr[root] {
		root = right
	}

	if root != idx {
		arr[idx], arr[root] = arr[root], arr[idx]

		heapify(arr, n, root)
	}
}

func buildHeap(arr []int, n int) {
	lastNode := n - 1

	for i := (lastNode - 1) / 2; i >= 0; i-- {
		heapify(arr, n, i)
	}
}
