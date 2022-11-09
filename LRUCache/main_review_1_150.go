package main

import "fmt"

func main() {
	lru := Constructor(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)

	fmt.Println(lru.Get(1))
}

type LRUCache struct {
	size  int
	mp    map[int]*Node
	count int
	head  *Node
	tail  *Node
}

type Node struct {
	Val  int
	Key  int
	Last *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	tail.Last = head
	head.Next = tail
	return LRUCache{size: capacity, tail: tail, head: head, mp: map[int]*Node{}}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.mp[key]
	if !ok {
		return -1
	}
	moveToHead(this.head, v)
	return v.Val
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.mp[key]; ok {
		v.Val = value
		moveToHead(this.head, v)
	} else if this.count >= this.size {
		removed := removeTail(this.tail)
		delete(this.mp, removed.Key)
		node := &Node{
			Val: value,
			Key: key,
		}
		addToHead(this.head, node)
		this.mp[key] = node
	} else {
		this.mp[key] = &Node{
			Key: key,
			Val: value,
		}
		addToHead(this.head, this.mp[key])
		this.count++
	}

}

func addToHead(head *Node, n *Node) {
	n.Next = head.Next
	head.Next.Last = n
	head.Next = n
	n.Last = head
}
func removeTail(tail *Node) *Node {
	last := tail.Last
	tail.Last = last.Last
	tail.Last.Next = tail
	return last
}
func moveToHead(head *Node, node *Node) {
	node.Last.Next = node.Next
	node.Next.Last = node.Last
	addToHead(head, node)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
