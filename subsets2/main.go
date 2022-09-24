package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{
		1, 2, 2,
	}

	fmt.Println(subsetsWithDup2(a))

}

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})
	var placeNum func(c, rem []int)
	placeNum = func(curr []int, rem []int) {

		item := make([]int, len(curr))
		copy(item, curr)
		result = append(result, item)
		if len(rem) == 0 {
			return
		}
		for index := range rem {
			if index == 0 || rem[index] != rem[index-1] {
				curr = append(curr, rem[index])
				placeNum(curr, rem[index+1:])
				curr = curr[:len(curr)-1]
			}

		}
	}
	placeNum([]int{}, nums)
	return result
}

func subsetsWithDup2(nums []int) [][]int {
    sort.Slice(nums, func (a,b int) bool{
        if nums[a] <= nums[b] {
            return true
        }
        return false
    })
    result := [][]int{}
    placeNum(&result, nums, []int{}, 0)
    return result
}

func placeNum(result *[][]int, nums[]int, cur []int, index int) {
    item := make([]int, len(cur))
    copy(item, cur)
    *result = append(*result, item)
    if index >= len(nums) {
        return
    }
    
    for i:=index;i<len(nums);i++ {
        if i>index && nums[i] == nums[i-1] {
            continue
        }
        cur = append(cur, nums[i])
        placeNum(result, nums, cur, i+1)
        cur = cur[:len(cur)-1]
    }
} 