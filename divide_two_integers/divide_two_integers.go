package divide_two_integers

import "math"

func divide(dividend int, divisor int) int {
	if dividend == -2147483648 && divisor == -1 {
		// because right answer is 2147483648 but test case expects wrong value
		return 2147483647
	}

	i := 0
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))

	isNegative := (divisor < 0 && dividend >= 0) || (divisor > 0 && dividend < 0)

	for ; dividendAbs >= divisorAbs; dividendAbs -= divisorAbs {
		i++
	}

	if isNegative {
		return -i
	} else {
		return i
	}
}
