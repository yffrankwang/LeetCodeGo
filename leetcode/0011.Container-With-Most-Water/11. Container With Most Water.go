package leetcode

func maxArea(height []int) int {
	max, l, r := 0, 0, len(height)-1
	for l < r {
		w := r - l // width
		h := 0     // height
		if height[l] < height[r] {
			h = height[l]
			l++
		} else {
			h = height[r]
			r--
		}

		temp := w * h
		if temp > max {
			max = temp
		}
	}
	return max
}
