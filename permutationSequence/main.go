package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getPermutation(4, 1))
	fmt.Println(getPermutation(4, 2))
	fmt.Println(getPermutation(4, 3))
	fmt.Println(getPermutation(4, 4))
	fmt.Println(getPermutation(4, 5))
	fmt.Println(getPermutation(4, 6))
}

// 根据n-1! 确定n个数有多少个组合
// 根据k确定初始 1 或者2或者3,4, 依次再确定每个组下面的组合直到最后
// 1       2      	3   	4
// 1 234   2 134
// 1 243   2 143
// ...
func getPermutation(n int, k int) string {
	nums := make([]int, n)
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		nums[i-1] = i
		fact[i] = i * fact[i-1]
	}
	s := []string{}

	newNums := nums
	var temp []int
	k--
	for len(newNums) > 0 {
		block := k / fact[n-1]
		s = append(s, fmt.Sprint(newNums[block]))
		temp = []int{}
		temp = append(temp, newNums[0:block]...)
		temp = append(temp, newNums[block+1:]...)
		k = k % fact[n-1]
		newNums = temp
		n -= 1
	}
	// for _, v := range newNums {
	// 	s = append(s, fmt.Sprint(v))
	// }
	return strings.Join(s, "")
}

func getPermutation2(n int, k int) string {
	rest := make([]int, n)
	for i := 1; i <= n; i++ {
		rest[i-1] = i
	}
	result := []string{}

	place(&result, []string{}, rest, k)
	return result[k-1]
}

func place(result *[]string, current []string, rest []int, k int) {
	if len(rest) == 0 {

		*result = append(*result, strings.Join(current, ""))
		return
	}

	for index, v := range rest {
		newCurrent := make([]string, len(current))
		copy(newCurrent, current)
		newCurrent = append(newCurrent, fmt.Sprint(v))
		var newRest []int
		newRest = append(newRest, rest[0:index]...)
		newRest = append(newRest, rest[index+1:]...)
		place(result, newCurrent, newRest, k)
		if len(*result) == k {
			return
		}
	}

}
