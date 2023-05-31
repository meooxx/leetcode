package main

type NumArray struct {
	st *SegmentTree
}

type SegmentTree struct {
	Left  *SegmentTree
	Right *SegmentTree
	Sum   int
	Start int
	End   int
}

func buildTree(start, end int, nums []int) *SegmentTree {
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
		leftTree := buildTree(start, mid, nums)
		rightTree := buildTree(mid+1, end, nums)
		st.Left = leftTree
		st.Right = rightTree
		st.Sum = st.Left.Sum + st.Right.Sum
		return st
	}

}

func updateTree(i, val int, st *SegmentTree) {
	if i == st.Start && i == st.End {
		st.Sum = val
		return
	}
	mid := (st.End-st.Start)/2 + st.Start
	if i <= mid {
		updateTree(i, val, st.Left)
	} else {
		updateTree(i, val, st.Right)
	}
	st.Sum = st.Left.Sum + st.Right.Sum
}

func Constructor(nums []int) NumArray {
	return NumArray{
		st: buildTree(0, len(nums)-1, nums),
	}
}

func (this *NumArray) Update(index int, val int) {
	updateTree(index, val, this.st)
}

//	1-4   5-7
//
// 1-2  3
func (this *NumArray) SumRange(left int, right int) int {
	return sumRange(this.st, left, right)
}

func sumRange(st *SegmentTree, left, right int) int {
	if left == st.Start && right == st.End {
		return st.Sum
	}
	mid := (st.End-st.Start)/2 + st.Start
	if left > mid {
		return sumRange(st.Right, left, right)
	}
	if right <= mid {
		return sumRange(st.Left, left, right)
	}
	return sumRange(st.Left, left, mid) + sumRange(st.Right, mid+1, right)

}

/**
* Your NumArray object will be instantiated and called as such:
* obj := Constructor(nums);
* obj.Update(index,val);
* param_2 := obj.SumRange(left,right);
 */
