package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8, 10, 11, 22, 33, 6, 1}))
}

func maxCoins(nums []int) int {
	numPadding := make([]int, 0, len(nums)+2)
	numPadding = append(numPadding, 1)
	numPadding = append(numPadding, nums...)
	numPadding = append(numPadding, 1)

	n := len(numPadding)
	dp := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = 0
			} else {
				for k := i + 1; k < j; k++ {
					score := numPadding[i]*numPadding[k]*numPadding[j] + dp[i][k] + dp[k][j]
					if dp[i][j] < score {
						dp[i][j] = score
					}
				}
			}
		}
	}
	return dp[0][n-1]
}
