package best_time_to_buy_and_sell_stock

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		output int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			prices: []int{2, 4, 1},
			output: 2,
		},
		{
			prices: []int{2, 1, 2, 1, 0, 1, 2},
			output: 2,
		},
		{
			prices: []int{2, 1, 2, 0, 1},
			output: 1,
		},
	}

	for i, tcase := range cases {
		fmt.Printf("case %d\n", i)

		got := maxProfit(tcase.prices)
		if got != tcase.output {
			t.Fatalf("want %d, got %d, case %d", tcase.output, got, i)
		}
	}
}
