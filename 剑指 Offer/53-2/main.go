package main

import "fmt"

func main() {
	fmt.Println([]int{0, 1, 3})
	fmt.Println([]int{0, 1})
	fmt.Println([]int{1, 2})
}

func missingNumber(nums []int) int {
	len := len(nums)
	low := 0
	high := len - 1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] != mid {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low == nums[low] {
		return low + 1
	}
	return low
}
