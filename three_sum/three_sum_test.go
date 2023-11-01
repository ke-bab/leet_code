package three_sum

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums []int
		want [][]int
	}{
		//{
		//	nums: []int{-1, 0, 1, 2, -1, -4},
		//	want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		//},
		//{
		//	nums: []int{0, 1, 1},
		//	want: [][]int{},
		//},
		//{
		//	nums: []int{0, 0, 0},
		//	want: [][]int{{0, 0, 0}},
		//},
		//{
		//	nums: []int{0, 0, 0, 0},
		//	want: [][]int{{0, 0, 0}},
		//},
		{
			nums: []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
	}

	for i, tc := range testCases {
		got := threeSum(tc.nums)
		if !isOk(tc.want, got) {
			t.Fatalf("want: %v, got: %v, case: %d", tc.want, got, i)
		}
	}
}

func isOk(want [][]int, got [][]int) bool {
	if len(want) != len(got) {
		return false
	}
	for i := 0; i < len(got); i++ {
		found := false
		for j := 0; j < len(want); j++ {
			if threeIsOk(want[j], got[i]) {
				found = true
				break
			}
		}
		if found == false {
			return false
		}
	}

	return true
}

func threeIsOk(want []int, got []int) bool {
	if len(want) != len(got) {
		return false
	}
	found := false
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(want); j++ {
			if got[i] == want[j] {
				found = true
				break
			}
		}
		if found == false {
			return false
		}
	}

	return true
}
