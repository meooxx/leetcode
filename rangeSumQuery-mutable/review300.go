package main

type SegmentTree struct {
	Start int
	End   int
	Left  *SegmentTree
	Right *SegmentTree
	Sum   int
}

type NumArray struct {
	st *SegmentTree
}

func Constructor(nums []int) NumArray {

	st := buildTree(nums, 0, len(nums)-1)
	return NumArray{
		st,
	}
}

func buildTree(nums []int, start, end int) *SegmentTree {

	if start > end {
		return nil
	}
	st := &SegmentTree{
		Start: start,
		End:   end,
	}

	if start == end {
		st.Sum = nums[start]
		return st
	} else {

		mid := (end-start)/2 + start
		left := buildTree(nums, start, mid)
		right := buildTree(nums, mid+1, start)

		st.Left = left
		st.Right = right
		st.Sum = st.Left.Sum + st.Right.Sum
	}
	return st
}

func (this *NumArray) Update(index int, val int) {
	UpdateTree(*this.st, index, val)
}
func UpdateTree(st SegmentTree, index int, val int) {
	if st.Start == st.End && st.End == index {
		st.Sum = val
		return
	}

	mid := (st.End - st.Start) / 2
	if index <= mid {
		UpdateTree(*st.Left, index, val)
	} else {
		UpdateTree(*st.Right, index, val)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return sumRange(*this.st, left, right)
}

//	0 -  3
//
// 01, 12 3
func sumRange(st SegmentTree, left, right int) int {
	if left == st.Start && right == st.End {
		return st.Sum
	}
	mid := (-st.Start+st.End)/2 + st.Start

	if left > mid {
		return sumRange(*st.Right, left, right)

	}
	if right <= mid {
		return sumRange(*st.Left, left, right)
	}
	return sumRange(*st.Left, left, mid) + sumRange(*st.Right, mid+1, right)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
