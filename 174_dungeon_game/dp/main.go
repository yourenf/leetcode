package main

func main() {
	d := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{10, 30, -5},
	}
	println(calculateMinimumHP(d))
}

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	dp := make([][]int, m)

	for i := m - 1; i >= 0; i-- {
		dp[i] = make([]int, n)
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = dungeon[i][j]
			} else if i == m-1 {
				dp[i][j] = min(0, dp[i][j+1]+dungeon[i][j])
			} else if j == n-1 {
				dp[i][j] = min(0, dp[i+1][j]+dungeon[i][j])
			} else {
				hp := max(dp[i+1][j], dp[i][j+1]) + dungeon[i][j]
				dp[i][j] = min(0, hp)
			}
		}
	}
	if dp[0][0] >= 0 {
		return 1
	}
	return 0 - dp[0][0] + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
