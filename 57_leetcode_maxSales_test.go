package test

/*
	-2 1 -3 4 -1 2 1 -5 4
	0  1  2 3  4 5 6  7 8

0  -2 -1 -4  0 -1 1 2 -3 1
1	0  1 -2  2  1 3 4 -1 3
2
3
4
5
6
7
8
*/
func maxSales(sales []int) int {
	var max int
	length := len(sales)
	dp := make([]int, length)
	max, dp[0] = sales[0], sales[0]
	for i := 1; i < length; i++ {
		if dp[i-1] < 0 {
			dp[i] = sales[i]
		} else {
			dp[i] = sales[i] + dp[i-1]
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
