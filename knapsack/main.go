package main

import "fmt"

/**
 * 动态规划求解0-1背包问题
 * @param w 物品重量
 * @param v 物品价值
 * @param c 背包容量
 * @return 背包能装下的最大价值
 * 时间复杂度：O(n*c)
 * 空间复杂度：O(n*c)
 * 优点：时间复杂度较低，能够求解较大规模的问题
 * 缺点：空间复杂度较高，不适用于数据量较大的问题
 */
func knapsack(w []int, v []int, c int) int {
	//物品数量
	n := len(w)
	//
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, c+1)
	}
	//将dp[0]作为原始缓存，从dp[1]开始装
	for i := 1; i <= n; i++ {
		for j := 1; j <= c; j++ {
			if j < w[i-1] {
				// 背包容量不足，不能装入第i个物品
				dp[i][j] = dp[i-1][j]
			} else {
				// 能装入第i个物品
				s := dp[i-1][j]
				t := dp[i-1][j-w[i-1]]+v[i-1]
				fmt.Println(s, t)
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i-1]]+v[i-1])
			}
		}
	}
	// 返回背包能装下的最大价值
	return dp[n][c]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(knapsack([]int{6, 7, 8}, []int{10, 9, 4}, 15))
}
