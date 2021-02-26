package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 给定一个数组，随机生成与一棵二叉检索树
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	node := build(0, len(nums)-1, nums)

	printTree(node)
	println()
	printTree2(node)
}

func build(start, end int, nums []int) *TreeNode {
	// 取第一个点，将左边的生成一棵树，右边的生成一棵树
	if start > end {
		return nil
	}
	middle := start + (end-start)/2
	node := &TreeNode{Val: nums[middle]}

	node.Left = build(start, middle-1, nums)
	node.Right = build(middle+1, end, nums)

	return node
}

func printTree(node *TreeNode) {
	fmt.Printf("%d\t", node.Val)

	if node.Left != nil {
		printTree(node.Left)
	}

	if node.Right != nil {
		printTree(node.Right)
	}
}

func printTree2(node *TreeNode) {
	if node.Left != nil {
		printTree(node.Left)
	}

	fmt.Printf("%d\t", node.Val)

	if node.Right != nil {
		printTree(node.Right)
	}
}
