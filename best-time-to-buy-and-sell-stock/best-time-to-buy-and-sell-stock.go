package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	min := -1
	profit := 0

	if len(prices) == 1 {
		return 0
	}

	for i := 0; i < len(prices); i++ {
		if min == -1 {
			min = prices[i]
		} else {
			if prices[i] < min {
				min = prices[i]
			}
		}
		if profit < prices[i]-min {
			profit = prices[i] - min
		}
	}
	if min == -1 {
		return 0
	}

	return profit
}
