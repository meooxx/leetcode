package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(largestNumber([]int{
		33,
		332,
		34,
		330,
	}))
	fmt.Println(largestNumber([]int{
		// 432 43243
		// 432 end从头开始 43         2  == 432 43243 == "432" + "43243"
		// 432            43 (也从头) 4  == 43243 4  == "43243" + "432"
		// 其实也是比较 a+b和b+a俩中情况
		432,
		43243,
	}))
	fmt.Println(largestNumber([]int{
		55,
		555,
	}))
	fmt.Println(largestNumber([]int{
		0,
		0,
	}))
}

func largestNumber1(nums []int) string {
	strs := []string{}

	for i := range nums {
		strs = append(strs, fmt.Sprint(nums[i]))
	}
	sort.Slice(strs, func(a, b int) bool {
		strA := strs[a]
		strB := strs[b]
		i, j := 0, 0
		for i < len(strA) || j < len(strB) {
			b1, b2 := strA[0], strB[0]
			if i < len(strA) {
				b1 = strA[i]
			} else {
				i = 0
			}
			if j < len(strB) {
				b2 = strB[j]
			} else {
				j = 0
			}
			if b1 < b2 {
				return false
			} else if b1 > b2 {
				return true
			}
			i++
			j++
		}

		// if j < len(strB) {
		//     return true
		// }
		return true
	})
	// for left, right := 0, len(strs)-1; left < right; {
	// 	strs[left], strs[right] = strs[right], strs[left]
	// 	left++
	// 	right--
	// }
	for i := range strs {
		if strs[i] != "0" {
			break
		}
		return "0"
	}
	return strings.Join(strs, "")
}

// 330, 331
// "330" + "331" compare 331+"330"
func largestNumber(nums []int) string {
	arr := make([]string, len(nums))
	for i := range nums {
		arr[i] = fmt.Sprint(nums[i])
	}
	sort.Slice(arr, func(a, b int) bool {
		sum1 := arr[a] + arr[b]
		sum2 := arr[b] + arr[a]
		return sum1 >= sum2
	})
	// because arr is a descending array
	// if arr[0] == 0  ==> every element of arr is 0
	if arr[0] == "0" {
		return "0"
	}
	return strings.Join(arr, "")
}
