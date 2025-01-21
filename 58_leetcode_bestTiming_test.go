package test

func bestTiming(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyDp := make([]int, len(prices))
	saleDp := make([]int, len(prices))
	length := len(prices)
	buyDp[0], saleDp[length-1] = prices[0], prices[length-1]
	for i := 1; i < length; i++ {
		if prices[i] < buyDp[i-1] {
			buyDp[i] = prices[i]
		} else {
			buyDp[i] = buyDp[i-1]
		}
		j := length - 1 - i
		if prices[j] > saleDp[j+1] {
			saleDp[j] = prices[j]
		} else {
			saleDp[j] = saleDp[j+1]
		}
	}
	max := saleDp[0] - buyDp[0]
	for i := 1; i < length; i++ {
		diff := saleDp[i] - buyDp[i]
		if diff > max {
			max = diff
		}
	}
	return max
}
