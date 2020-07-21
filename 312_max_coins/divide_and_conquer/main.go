package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8, 10, 11, 22, 33, 6, 1}))
}

func maxCoins(nums []int) int {
	padding := make([]int, 0, len(nums)+2)
	padding = append(padding, 1)
	padding = append(padding, nums...)
	padding = append(padding, 1)
	n := len(padding)
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}
	score := helper(padding, 0, n-1, memo)
	fmt.Println(memo)
	return score
}

func helper(nums []int, begin int, end int, memo [][]int) int {
	if begin+1 == end {
		return 0
	}
	if memo[begin][end] > 0 {
		return memo[begin][end]
	}
	max := 0
	for i := begin + 1; i < end; i++ {
		score := nums[i]*nums[begin]*nums[end] + helper(nums, begin, i, memo) + helper(nums, i, end, memo)
		if score > max {
			max = score
		}
	}
	memo[begin][end] = max
	return max
}
