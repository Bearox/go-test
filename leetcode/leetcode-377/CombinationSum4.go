package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(combinationSum4(nums, 32))
}

func combinationSum4(nums []int, target int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	return combinationSumForSortedNums(nums, target)
}

// nums是从大到小排序的数组
func combinationSumForSortedNums(nums[] int, target int) int  {
	// 数组空就是0
	if nil == nums || len(nums) == 0 {
		return 0
	}

	// 最小的元素就比目标值大，那就无解了
	if nums[0] > target {
		return 0
	}

	dp := make([]int, target + 1)
	dp[0] = 1
	for idx := 1; idx <= target; idx++ {
		dp[idx] = 0
		for _, elem := range nums {
			if elem > idx {
				break
			}
			dp[idx] += dp[idx - elem]
		}
	}
	return dp[target]
}
