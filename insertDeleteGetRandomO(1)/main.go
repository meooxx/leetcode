package main
import "math/rand"

type RandomizedSet struct {
	list []int
	set  map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		list: []int{},
		set:  map[int]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.list = append(this.list, val)
	this.set[val] = len(this.list) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if loc, ok := this.set[val]; ok {
		if loc < len(this.list)-1 {
			lastone := this.list[len(this.list)-1]
			this.list[loc] = lastone
			this.set[lastone] = loc
		}
		delete(this.set, val)
		this.list = this.list[:len(this.list)-1]
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(len(this.list))]
}

/**
* Your RandomizedSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Insert(val);
* param_2 := obj.Remove(val);
* param_3 := obj.GetRandom();
 */
