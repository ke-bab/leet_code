package search_insert_position

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			want:   2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			want:   1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			want:   4,
		},
	}

	for i, tc := range testCases {
		got := searchInsert(tc.nums, tc.target)
		if got != tc.want {
			t.Fatalf("want: %d, got: %d, case: %d", tc.want, got, i)
		}
	}
}
