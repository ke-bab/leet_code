package valid_palindrome

import "testing"

func TestValidPalindrome(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{
			input: "A man, a plan, a canal: Panama",
			want:  true,
		},
		{
			input: "race a car",
			want:  false,
		},
	}

	for tci, tc := range testCases {
		got := isPalindrome(tc.input)
		if tc.want != got {
			t.Fatalf("want len %t != got len %t, case %d", tc.want, got, tci)
		}
	}
}
