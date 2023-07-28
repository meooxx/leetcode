package main

import "sort"

type SummaryRanges struct {
	mp map[int]bool
}

func Constructor() SummaryRanges {
	return SummaryRanges{map[int]bool{}}
}

func (this *SummaryRanges) AddNum(value int) {
	this.mp[value] = true
}

func (this *SummaryRanges) GetIntervals() [][]int {
	keys := []int{}
	for i := range this.mp {
		keys = append(keys, i)
	}
	sort.Slice(keys, func(a, b int) bool {
		return keys[a] < keys[b]
	})

	result := [][]int{}

	for i := range keys {
		if len(result) > 0 && result[len(result)-1][1]+1 == keys[i] {
			result[len(result)-1][1] = keys[i]
		} else {
			result = append(result, []int{keys[i], keys[i]})
		}
	}
	return result
}

/**
* Your SummaryRanges object will be instantiated and called as such:
* obj := Constructor();
* obj.AddNum(value);
* param_2 := obj.GetIntervals();
 */
