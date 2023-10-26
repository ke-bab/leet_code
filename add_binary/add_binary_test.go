package add_binary

import "testing"

func TestAddBinary(t *testing.T) {
	testCases := []struct {
		inputA string
		inputB string
		want   string
	}{
		{
			inputA: "11",
			inputB: "1",
			want:   "100",
		},
	}

	for tci, tc := range testCases {
		res := addBinary(tc.inputA, tc.inputB)
		if res != tc.want {
			t.Fatalf("res %s not equal want %s, for case %d", res, tc.want, tci)
		}
	}
}
