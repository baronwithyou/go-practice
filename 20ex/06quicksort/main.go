package main

import "fmt"

func main() {
	arr := []int{9, 4, 2, 8, 7, 1}

	sort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func sort(arr []int, l, r int) {
	if l < r {
		pivot := partition(arr, l, r)

		sort(arr, l, pivot-1)
		sort(arr, pivot+1, r)
	}
}

// 为某元素指定位置
func partition(arr []int, l, r int) int {
	pivot := arr[l]

	for l < r {
		for l < r && arr[l] < pivot {
			l++
		}

		for l < r && arr[r] > pivot {
			r--
		}

		arr[l], arr[r] = arr[r], arr[l]
	}

	arr[l] = pivot
	return l
}
