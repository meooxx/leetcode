package main

import (
	"fmt"
	_ "strconv"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	r := []int{}
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			r = append(r, nums1[i])
			i++
		} else {
			r = append(r, nums2[j])
			j++
		}
	}
	r = append(r, nums1[i:]...)
	r = append(r, nums2[j:]...)
	if len(r)%2 == 1 {
		return float64(r[(len(r)-1)/2])
	} else {
		mid := len(r) / 2
		num := (float64(r[mid-1] + r[mid])) / 2
		return num
	}

}

func main() {
	v := findMedianSortedArrays([]int{1, 1}, []int{1, 2})
	v1 := findMedianSortedArrays([]int{1,2}, []int{3, 4})

	fmt.Println(v, v1)

}
