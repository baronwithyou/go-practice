package main

import "fmt"

const VERTICES = 6

func initiate(parent []int) {
	for i := range parent {
		parent[i] = -1
	}
}

func find(x int, parent []int) int {
	if parent[x] == -1 {
		return x
	}

	return find(parent[x], parent)
}

func merge(x, y int, parent []int) bool {
	xParent := find(x, parent)
	yParent := find(y, parent)

	if xParent == yParent {
		return false
	} else {
		parent[xParent] = yParent

		return true
	}
}

func main() {
	parent := make([]int, VERTICES)
	initiate(parent)

	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 4}, {2, 5}}

	for _, edge := range edges {
		x := edge[0]
		y := edge[1]

		if !merge(x, y, parent) {
			fmt.Println("Cycle detected!")
			return
		}
	}
	fmt.Println("Not found cycle")
}
