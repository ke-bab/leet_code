package search_insert_position

func searchInsert(nums []int, target int) int {
	result := -1
	indexToInsert := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
			break
		}
		if nums[i] < target {
			indexToInsert = i + 1
		}
	}
	if result == -1 {
		return indexToInsert
	}

	return result
}
