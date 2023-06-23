package main

import "fmt"
//  3 		2    5 1 
//  [3] [3,2] [5] [5,1]   maxArr
//            pop all elements are less than 5
// 
// maxArr1(0), maxArr1(1) maxArr1(2)..
// maxArr2(5)  maxArr2(4) maxArr2(3)...
///   k== 5 
//  
// merge maxArr1 and maxArr2
// and find the most max number arr

func main() {

	// fmt.Println(maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	// fmt.Println(maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))

	// fmt.Println(maxNumber([]int{4,9,5}, []int{8,7,4}, 3))
	fmt.Println(maxNumber([]int{2, 5, 6, 4, 4, 0}, []int{7, 3, 8, 0, 6, 5, 7, 6, 2}, 15))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	answer := make([]int, k)
	for i := max(0, k-len(nums2)); i <= len(nums1) && i <= k; i++ {
		a1 := maxArr(nums1, i)
		a2 := maxArr(nums2, k-i)
		merged := merge(a1, a2, k)
		for p := 0; p < k; {
			if merged[p] == answer[p] {
				p++
			} else {
				if merged[p] > answer[p] {
					answer = merged
				}

				break
			}
		}
	}
	return answer
}

func greater(num1 []int, i int, num2 []int, j int) bool {
	for i < len(num1) && j < len(num2) && num1[i] == num2[j] {
		i++
		j++
	}
	return j == len(num2) || (i < len(num1) && num1[i] > num2[j])
}

func merge(num1 []int, num2 []int, k int) []int {
	anwser := make([]int, k)
	// 603
	// 9820   9 8 6 2 0 3 (0)
	//                     | 9820
	// 9:6 							9
	// 8:6              8
	// 0(top):2 				2
	// 0(top):0(bottom) 0 top, why 0 at top? 
	//   next of top 0 is 3 is great bottom
	for i, j, r := 0, 0, 0; r < k; r++ {
		if greater(num1, i, num2, j) {
			anwser[r] = num1[i]
			i++
		} else {
			anwser[r] = num2[j]
			j++
		}
	}
	return anwser
}

func maxArr(nums []int, n int) []int {
	stack := make([]int, n)
	size := 0

	for i := 0; i < len(nums); i++ {
		//   size
		// 5 3,     8
		for len(nums)-i > n-size && size > 0 && nums[i] > stack[size-1] {
			size--
		}
		if size < n {
			stack[size] = nums[i]
			size++
		}
	}
	return stack
}
