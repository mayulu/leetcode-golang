package leetcode_go_mayulu

func TwoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {

				result = append(result, i)
				result = append(result, j)
				return result
			}

		}
	}
	return res
}
