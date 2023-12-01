package majority_element

func majorityElement(nums []int) int {

	m := 0
	val := 0
	counts := map[int]int{}
	for i := 0; i < len(nums); i++ {
		counts[nums[i]]++
	}
	for k, v := range counts {
		if v > m {
			val = k
			m = v
		}
	}

	return val
}
