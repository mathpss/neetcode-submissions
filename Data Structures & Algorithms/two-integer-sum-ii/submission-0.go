func twoSum(numbers []int, target int) []int {
	myMap := make(map[int]int)

	for idx, valor := range numbers {
		complement := target - valor
		_, ok := myMap[complement]

		if ok {
			return []int{myMap[complement], idx +1}
		} else {
			myMap[valor] = idx +1
		}
	}
	return []int{}
}
