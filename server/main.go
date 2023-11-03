package main

import "fmt"

func adjustArray(nums []int) []int {
	// 使用双指针法，一个指针从头开始，一个指针从尾部开始
	left := 0
	right := len(nums) - 1

	for left < right {
		// 如果左指针指向的是偶数，右指针指向的是奇数，则交换两个数字的位置
		if nums[left]%2 == 0 && nums[right]%2 == 1 {
			nums[left], nums[right] = nums[right], nums[left]
		}

		// 左指针移动到下一个位置，直到指向奇数
		if nums[left]%2 == 1 {
			left++
		}

		// 右指针移动到下一个位置，直到指向偶数
		if nums[right]%2 == 0 {
			right--
		}
	}

	return nums
}

// 01背包
// w: 物品重量
// v: 物品价值
// c: 背包容量
func knapsack(w []int, v []int, c int) int {
	// dp定义：dp[i][j]表示放入索引为i的物品，背包容量为j时，背包中物品的最大价值
	dp := make([][]int, len(w))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, c+1)
	}
	// 初始化dp
	// 将第0号物品放入背包
	for cc := 0; cc <= c; cc++ {
		if cc >= w[0] {
			dp[0][cc] = v[0]
		}
	}

	// 状态转移
	for i := 1; i < len(w); i++ {
		for cc := 0; cc <= c; cc++ {
			count := 0
			// 能放入背包，count=1
			if cc >= w[i] {
				count = 1
			}
			for k := 0; k <= count; k++ {
				// 不放入背包 就是上一个物品的最大价值
				// dp[i][cc] = dp[i-1][cc]
				// 放入背包 就是上一个物品的最大价值加上当前物品的价值，但是要减去当前物品的重量来放入背包
				// dp[i][cc] = v[i] + dp[i-1][cc-w[i]]
				dp[i][cc] = Max(dp[i-1][cc], k*v[i]+dp[i-1][cc-k*w[i]])
			}
		}
	}
	// 打印dp数组
	for i := range dp {
		fmt.Print(dp[i], " ")
		fmt.Println()
	}

	return dp[len(w)-1][c]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapsack1(w []int, v []int, c int) int {
	// dp定义：dp[i][j]表示放入索引为i的物品，背包容量为j时，背包中物品的最大价值
	dp := make([][]int, len(w))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, c+1)
	}
	// 初始化dp
	// 将第0号物品放入背包
	for cc := 0; cc <= c; cc++ {
		// 可放入多次
		dp[0][cc] = (cc / w[0]) * v[0]
	}

	// 状态转移
	for i := 1; i < len(w); i++ {
		for cc := 0; cc <= c; cc++ {
			// 当前背包容量可以放入的最大物品数量
			count := cc / w[i]
			for k := 0; k <= count; k++ {
				// 不放入背包 就是上一个物品的最大价值
				// dp[i][cc] = dp[i-1][cc]
				// 放入多次背包 就是上一个物品的最大价值加上当前物品的价值，但是要减去当前物品的重量来放入背包
				// dp[i][cc] = k*v[i] + dp[i-1][cc-k*w[i]]
				// 当 k=0时, k*v[i]+dp[i-1][cc-k*w[i]] 也就是 dp[i-1][cc]
				dp[i][cc] = Max(dp[i][cc], k*v[i]+dp[i-1][cc-k*w[i]])
			}
		}
	}

	// 打印dp数组
	for i := range dp {
		fmt.Print(dp[i], " ")
		fmt.Println()
	}

	return dp[len(w)-1][c]
}

func main() {
	//context.Background()
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println("Original array:", nums)
	//fmt.Println("Adjusted array:", adjustArray(nums))
	//fmt.Println(knapsack([]int{3, 4, 2, 5}, []int{8, 14, 6, 12}, 10))
	fmt.Println(knapsack1([]int{3, 4, 2, 5}, []int{8, 14, 6, 12}, 10))
}
