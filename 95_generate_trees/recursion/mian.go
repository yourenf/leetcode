package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	list := generateTrees(3)
	for _, val := range list {
		fmt.Println(val)
	}
}

func generateTrees(n int) []*TreeNode {
	if n > 0 {
		return helper(1, n)
	}
	return []*TreeNode{}
}

func helper(begin, end int) []*TreeNode {
	result := make([]*TreeNode, 0)
	if begin > end {
		result = append(result, nil)
		return result
	}
	for i := begin; i <= end; i++ {
		left := helper(begin, i-1)
		right := helper(i+1, end)
		for l := 0; l < len(left); l++ {
			for r := 0; r < len(right); r++ {
				node := &TreeNode{Val: i, Left: left[l], Right: right[r]}
				result = append(result, node)
			}
		}
	}
	return result
}
