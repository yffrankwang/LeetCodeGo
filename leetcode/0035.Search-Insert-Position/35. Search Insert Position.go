package leetcode

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		h := i + (j-i)>>1 // avoid overflow when computing h
		if nums[h] < target {
			i = h + 1
		} else {
			j = h
		}
	}

	return i
}
