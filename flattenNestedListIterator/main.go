package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */



type NestedIterator struct {
	stack []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	iter := &NestedIterator{[]int{}}
	helper(nestedList, iter)
	return iter
}

func (this *NestedIterator) Next() int {
	cur := this.stack[0]
	this.stack = this.stack[1:]
	return cur
}

func (this *NestedIterator) HasNext() bool {
	return len(this.stack) > 0
}

func helper(nestedList []*NestedInteger, iter *NestedIterator) {
	for i := 0; i < len(nestedList); i++ {
		if nestedList[i].IsInteger() {
			(*iter).stack = append((*iter).stack, nestedList[i].GetInteger())
		} else {
			helper(nestedList[i].GetList(), iter)
		}
	}

}
