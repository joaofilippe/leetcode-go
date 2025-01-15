package maxarea

func maxArea(height []int) int {
	maxArea := 0
	l, r := 0, len(height)-1

	for l < r {
		maxArea = max(maxArea, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}
