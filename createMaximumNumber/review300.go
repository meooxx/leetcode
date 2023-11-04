package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	anwser := make([]int, k)

	// incaseof  [1,2, 3, 4, 5] , [1],  4
	// i = 1, k-i = 3, and 3 > len([1])
	// k - len(nums2) > 0 ? e.g. 5 - 3,  nums from 2 to min(5, len(nums1)), and k-i from len(nums2) to 0
	// or i from 0 to min(5, min(nums1))
	for i := max(0, k-len(nums2)); i <= k && i <= len(nums1); i++ {
		max1 := getMaxArr(nums1, i)
		max2 := getMaxArr(nums2, k-i)
		maxVal := merge(max1, max2)
		// fmt.Println(max1, max2, maxVal)
		for p, q := 0, 0; p < len(maxVal); {
			if maxVal[p] > anwser[q] {
				anwser = maxVal
				break
			} else if maxVal[p] == anwser[q] {
				p++
				q++
			} else {
				break
			}
		}
	}
	return anwser
}

// 5 3 6
// 5 -> 5,3 -> 6 -> 56
func getMaxArr(nums []int, l int) []int {
	stack := make([]int, l)
	index := 0
	for i := 0; i < len(nums); i++ {
		for index > 0 && stack[index-1] < nums[i] && len(nums)-i > l-index {
			index--
		}
		if index < l {
			stack[index] = nums[i]
			index++
		}
	}
	return stack

}

// 1, 2
// 1, ”  ”, 1
// 12, 11
func greater(a1 []int, i int, a2 []int, j int) bool {

	for i < len(a1) && j < len(a2) && a1[i] == a2[j] {
		i++
		j++
	}
	return j >= len(a2) || (i < len(a1) && a1[i] > a2[j])

}

func merge(a1, a2 []int) []int {
	i, j := 0, 0
	result := make([]int, len(a1)+len(a2))
	index := 0
	for index < len(result) {
		if greater(a1, i, a2, j) {
			result[index] = a1[i]
			i++
		} else {
			result[index] = a2[j]
			j++
		}
		index++

	}
	return result
}
