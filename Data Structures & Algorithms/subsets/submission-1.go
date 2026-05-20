func subsets(nums []int, ) [][]int {
	var result [][]int
	var current []int

	backtrack(0, current, nums, &result)

	return result
}

func backtrack(index int, currentItem []int, nums []int, result *[][]int) {
	if index == len(nums) {
		temp := make([]int, len(currentItem))
		copy(temp,currentItem)
		*result = append(*result, temp)
		return
	}
	
	currentItem = append(currentItem, nums[index])
	backtrack(index+1, currentItem, nums, result)
	currentItem = currentItem[:len(currentItem)-1]
	backtrack(index+1, currentItem, nums, result)
}