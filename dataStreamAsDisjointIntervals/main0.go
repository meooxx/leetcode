package main

type Node struct {
	Val  int
	Next *Node
}

type SummaryRanges struct {
	List *Node
}

func Constructor() SummaryRanges {
	return SummaryRanges{&Node{
		Val: -1,
	}}
}

func (this *SummaryRanges) AddNum(value int) {
	cur := this.List

	for cur != nil {
		if cur.Next == nil {
			cur.Next = &Node{Val: value}
			return
		}
		if value < cur.Next.Val {
			cur.Next = &Node{
				Val:  value,
				Next: cur.Next,
			}
			break
		} else {
			cur = cur.Next
		}

	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	cur := this.List.Next
	if cur == nil {
		return [][]int{}
	}
	start := cur.Val
	end := cur.Val
	result := [][]int{}
	for cur != nil {
		if cur.Next != nil && cur.Next.Val <= cur.Val+1 {
			end = cur.Next.Val
			cur = cur.Next
		} else {
			result = append(result, []int{
				start, end,
			})
			if cur.Next != nil {
				start = cur.Next.Val
				end = cur.Next.Val
			}

			cur = cur.Next
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