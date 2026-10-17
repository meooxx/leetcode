package main

import "math/rand"

type RandomizedCollection struct {
	list [][]int
	mp   map[int][]int
}

func Constructor() RandomizedCollection {
	return RandomizedCollection{
		list: [][]int{},
		mp:   map[int][]int{},
	}
}

func (this *RandomizedCollection) Insert(val int) bool {
	v, ok := this.mp[val]
	if ok {
		this.mp[val] = append(v, len(this.list))
	} else {
		this.mp[val] = []int{len(this.list)}
	}
	this.list = append(this.list, []int{val, len(this.mp[val]) - 1})
	return !ok
}

func (this *RandomizedCollection) Remove(val int) bool {
	indices, ok := this.mp[val]
	if ok {
		// last of nums list
		last := this.list[len(this.list)-1]
		// From last to first
		// After removing, index in mp  will great than the lenght of list
		lastIndice := len(indices) - 1
		// swap the first index of val and last element in nums
		this.list[indices[lastIndice]] = last

		// update mp[last]
		this.mp[last[0]][last[1]] = indices[lastIndice]

		// remove the last element
		this.list = this.list[:len(this.list)-1]
		// remove  val's index from mp[val]
		this.mp[val] = indices[:lastIndice]
		if len(this.mp[val]) == 0 {
			delete(this.mp, val)
		}

	}
	return ok

}

func (this *RandomizedCollection) GetRandom() int {
	index := rand.Intn(len(this.list))
	return this.list[index][0]
}

/**
* Your RandomizedCollection object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Insert(val);
* param_2 := obj.Remove(val);
* param_3 := obj.GetRandom();
 */
