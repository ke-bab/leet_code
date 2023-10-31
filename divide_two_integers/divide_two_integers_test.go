package divide_two_integers

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},
		{
			dividend: -1,
			divisor:  1,
			want:     -1,
		},
		{
			dividend: -2147483648,
			divisor:  -1,
			want:     2147483647,
		},
	}

	for i, tc := range testCases {
		result := divide(tc.dividend, tc.divisor)
		if result != tc.want {
			t.Fatalf("result: %d is not equal wanted: %d, case: %d", result, tc.want, i)
		}
	}
}
