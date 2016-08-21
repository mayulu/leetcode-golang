package leetcode_mayulu

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target-nums[i] == nums[j] {

				res = append(res, i)
				res = append(res, j)

				return res
			}

		}
	}
	return res
}
