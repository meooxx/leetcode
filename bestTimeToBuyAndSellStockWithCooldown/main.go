package main

//	1 ,  2,  3,  0,   2
//
// s0   0     0    0   1   2    2
// s1  -1    -1   -1  -1   1    1
// s2  -Inf   0    1   2   1    3
func maxProfit(prices []int) int {
	s0 := make([]int, len(prices)+1)
	s1 := make([]int, len(prices)+1)
	s2 := make([]int, len(prices)+1)
	s0[0] = 0
	s1[0] = -prices[0]
	s2[0] = -10001
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := range prices {
		s0[i+1] = max(s0[i], s2[i])
		s1[i+1] = max(s1[i], s0[i]-prices[i])
		s2[i+1] = s1[i] + prices[i]
	}
	fmt.Println(s0, s2)
	return max(s0[len(prices)], s2[len(prices)])
}
