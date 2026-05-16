func threeSum(nums []int) [][]int {
	var result [][]int
	myMap := make(map[[3]int]struct{})
	i := 0

	for i < len(nums)-2 {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {

					if _, ok := myMap[[3]int{nums[i], nums[j], nums[k]}]; ok {
						continue
					}

					result = append(result, []int{nums[i], nums[j], nums[k]})
					myMap[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
					myMap[[3]int{nums[i], nums[k], nums[j]}] = struct{}{}
					myMap[[3]int{nums[k], nums[i], nums[j]}] = struct{}{}
					myMap[[3]int{nums[k], nums[j], nums[i]}] = struct{}{}
					myMap[[3]int{nums[j], nums[k], nums[i]}] = struct{}{}
					myMap[[3]int{nums[j], nums[i], nums[k]}] = struct{}{}
				}
			}
		}
		i++
	}
	return result
}
