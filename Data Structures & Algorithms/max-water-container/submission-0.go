func maxArea(heights []int) int {
	l, r := 0, len(heights)-1
	var result float64
	for l < len(heights) {
		for r > l {
			lowestHeights := math.Min(float64(heights[l]), float64(heights[r]))
			differeceBase := r - l

			result = math.Max(result, (lowestHeights * float64(differeceBase)))
			r--
		}
		r = len(heights)-1
		l++
	}

	return int(result)
}
