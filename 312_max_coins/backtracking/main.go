package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8, 10, 11, 22, 33, 6, 1}))
}

var max int = 0
var count int = 0

func maxCoins(nums []int) int {
	max = 0

	helper(nums, 0, 0)
	fmt.Println("count", count)
	return max
}

func helper(nums []int, level int, coins int) {
	if level == len(nums) {
		if coins > max {
			max = coins
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		count++
		curr := nums[i]
		if curr == -1 {
			continue
		}
		prev := getPrev(nums, i)
		next := getNext(nums, i)
		value := prev * curr * next
		nums[i] = -1
		helper(nums, level+1, coins+value)
		nums[i] = curr
	}
}

func getPrev(nums []int, i int) int {
	for j := i - 1; j >= 0; j-- {
		if nums[j] != -1 {
			return nums[j]
		}
	}
	return 1
}

func getNext(nums []int, i int) int {
	for j := i + 1; j < len(nums); j++ {
		if nums[j] != -1 {
			return nums[j]
		}
	}
	return 1
}
