package main
/*
*  1  2  3  4  5  6  8  10 12 ....
*  use a sub-sequence where every item is  multipied by 2,3,5
*  and apply merge sort to obtain the smallest one
*  min(1*2, 1*3, 1*5,) ...  min(10 * 2, 10*3, 10*5) ...
* 
*  1*2   2*2  3*2  4*2 5*2 6*2  8*2 ...
*  1*3   2*3  3*3  4*3 5*3 6*3  8*3 ...
*  1*5   2*5  3*5  4*5 5*5 6*5  8*5 ..
*/
func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	index1 := 0
	index2 := 0
	index3 := 0
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		v1 := nums[index1] * 2
		v2 := nums[index2] * 3
		v3 := nums[index3] * 5
		nums[i] = min(min(v1, v2), v3)
		if nums[i] == v1 {
			index1++
		}
		if nums[i] == v2 {
			index2++
		}
		if nums[i] == v3 {
			index3++
		}
	}
	return nums[n-1]
}
