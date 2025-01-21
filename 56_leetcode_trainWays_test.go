package test

func trainWays(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	a, b := 1, 1
	for i := 2; i <= num; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return b
}
