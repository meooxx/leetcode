package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,5,1,1,6,4}
	wiggleSort(nums)
	fmt.Println(nums)
}
// sort
func wiggleSort0(nums []int) {
	copied := make([]int, len(nums))
	copy(copied, nums)
	fmt.Println(copied)

	sort.Slice(copied, func(a, b int) bool {
		return copied[a] < copied[b]
	})
	fmt.Println(copied)
	index := len(nums) - 1
	for i := 1; i < len(nums); i += 2 {
		nums[i] = copied[index]
		index--
	}
	for i := 0; i < len(nums); i += 2 {
		nums[i] = copied[index]
		index--
	}
}
func getIndex(n int, l int) int {
	return (1 + 2*n) % (l | 1)
}
//  0 1 2 3 4 5
//  S L S L S L   s for small, l for large

// left  is next index that can put num > median                   
// right is next index that can put num < median
//                  (i)       			5 is a median of nums
//       nums[mapped(0)==1] == 6 >  5,  swap(mapped(left==0), 1) left++,i++ (left++ mean put 6) 
//       nums[mapped(1)==3] == 5 == 5,  i++
//       nums[mapped(2)==5] == 2 <  5,  swap(mapped(right==4), 5) right--(mean put 2) 	13   6    5    5   2  4
//       										== 4 <  5,  swap(mapped(right==2), 5) right-- 						  13   6    4    5   2  5
//													== 5 == 5,  i++
//            mapped(3)==0  == 13 > 5, 	swap(mapped(left==1), 0) left++, i++            5    6    4    13   2  5
//              left                    right
//               i    
// Original idx: 0    1    2    3    4    5  
// Mapped idx:   1    3    5    0    2    4 
// Array:        13   6    5    5    4    2 
// O(n) quick select
func wiggleSort(nums []int) {
	median := findKthLargestNum(nums, (len(nums)+1)/2)
	left, right := 0, len(nums)-1
	for i := 0; i <= right;{
		mappedIndex := getIndex(i, len(nums))
		if nums[mappedIndex] > median {
			mappedLeft := getIndex(left, len(nums))
			nums[mappedIndex], nums[mappedLeft] = nums[mappedLeft], nums[mappedIndex]
			left++
			// i++?
			// 7 10  0  8  3      6
			// i is great or equal left
			// and `nums[mappedIndex]` >= median
			// or the value of `mappedIndex` and `right` should be switched
			// so nums[mappedLeft] is always great or equal to  median
			i++
		} else if nums[mappedIndex] < median {
			mappedRight := getIndex(right, len(nums))
			nums[mappedIndex], nums[mappedRight] = nums[mappedRight], nums[mappedIndex]
			right--
		} else {
			i++
		}
	}
}

func findKthLargestNum(nums []int, n int) int {
	left, end := 0, len(nums)-1
	midVal := nums[end]
	right := end - 1
	// 1 7 4  5 =>  1 4 7 5
	for left <= right {
		for left <= right && nums[left] < midVal {
			left++
		}
		for left <= right && nums[right] >= midVal {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
	nums[end], nums[left] = nums[left], nums[end]
	size := end - left + 1
	if size == n {
		return nums[left]
	} else if size > n {
		return findKthLargestNum(nums[left+1:], n)
	} else {
		return findKthLargestNum(nums[:left], n-size)
	}
}
