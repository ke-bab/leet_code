package plus_one

import "testing"

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{1, 2, 3},
			want:  []int{1, 2, 4},
		},
		{
			input: []int{9},
			want:  []int{1, 0},
		},
	}

	for tci, tc := range testCases {

		got := plusOne(tc.input)
		if len(tc.want) != len(got) {
			t.Fatalf("want len %d != got len %d, case %d", len(tc.want), len(got), tci)
		}
		for i, v := range tc.want {
			if v != got[i] {
				t.Fatalf("want %d != got %d, case %d", tc.want, got, tci)
			}
		}
	}
}
