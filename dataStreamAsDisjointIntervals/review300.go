package main

import "sort"

type SummaryRanges struct {
	mp map[int]bool
}

func Constructor() SummaryRanges {
	return SummaryRanges{ map[int]bool{}}
}

func (this *SummaryRanges) AddNum(value int) {
	if !this.mp[value] {
		this.mp[value] = true
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	values := []int{}
	for v := range this.mp {
		values = append(values, v)
	}
	sort.Slice(values, func(a, b int) bool {
		return values[a] < values[b]
	})
	result := [][]int{}
	for i := range values {
		if len(result) > 0 && result[len(result)-1][1] + 1 == values[i] {
			result[len(result)-1][1] = values[i]

		}else {
			result = append(result, []int{values[i], values[i]})
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