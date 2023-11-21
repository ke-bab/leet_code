package count_and_say

import "testing"

func TestCountAndSay(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{
			input: 4,
			want:  "1211",
		},
	}

	for i, tc := range testCases {
		got := countAndSay(tc.input)
		if got != tc.want {
			t.Fatalf("want %s, got %s, case %d", tc.want, got, i)
		}
	}
}
