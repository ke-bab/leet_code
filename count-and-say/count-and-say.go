package count_and_say

import "strconv"

func countAndSay(n int) string {
	count := "1"
	for i := 1; i < n; i++ {
		count = countRecursive(count)
	}

	return count
}

func countRecursive(num string) string {
	res := ""
	count := 0
	prev := -1
	for _, v := range num {
		val, _ := strconv.Atoi(string(v))
		if prev == -1 {
			count++
			prev = val
			continue
		}
		if prev != val {
			res += strconv.Itoa(count) + strconv.Itoa(prev)
			count = 1
			prev = val
		} else {
			count++
			prev = val
		}
	}
	res += strconv.Itoa(count) + strconv.Itoa(prev)

	return res
}
