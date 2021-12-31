package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	last := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[last] {
			continue
		}
		last++
		if last != i {
			nums[last] = nums[i]
		}
	}
	return last + 1
}
