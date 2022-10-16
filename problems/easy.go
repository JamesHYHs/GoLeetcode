package problems

func Q14LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	mixPrefix := 0
	for {
		for i := 1; i < len(strs); i++ {
			if mixPrefix >= len(strs[i-1]) || mixPrefix >= len(strs[i]) || strs[i-1][mixPrefix] != strs[i][mixPrefix] {
				return strs[0][:mixPrefix]
			}
		}
		mixPrefix++
	}
	return strs[0][:mixPrefix]
}

func searchInsertQ35(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	if target > nums[right] {
		return len(nums)
	} else if target < nums[left] {
		return 0
	} else {
		for left < right {
			mid = (left + right) / 2
			if target == nums[mid] {
				return mid
			} else if target > nums[mid] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = (left + right) / 2
		if target > nums[mid] {
			return mid + 1
		} else {
			return mid
		}
	}
}

func findMaxAverageQ643(nums []int, k int) float64 {
	sum, start := 0, 0
	avg, maxAvg := 0.0, 0.0
	if len(nums) == 1 {
		return float64(nums[0])
	}
	for end := range nums {
		sum += nums[end]
		if end >= k-1 {
			avg = float64(sum) / float64(k)
			if avg > maxAvg || maxAvg == 0.0 {
				maxAvg = avg
			}
			sum -= nums[start]
			start++
		}
	}
	return maxAvg
}
