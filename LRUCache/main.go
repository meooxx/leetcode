package main

import "fmt"

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1), "Get 1")
	lru.Put(3, 3)
	fmt.Println(lru.Get(2), "Get 2")
	lru.Put(4, 4)
	fmt.Println(lru.Get(1), "Get 1")

	fmt.Println(lru.Get(3), "Get 3")
	fmt.Println(lru.Get(4), "Get 4")

}

// double link
// head and tail , head<->n1<->n2->n3<->tail
// if read/update a node, move the node to head
// if count > size, remove the tail.pre (eg.n3)
//   and insert the new node to head
type LRUCache struct {
	size   int
	values map[int]*Node
	count  int
	head   *Node
	tail   *Node
}

type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}

func addFirst(head *Node, node *Node) {
	// head a b
	// head node a b
	node.Next = head.Next
	node.Next.Pre = node
	node.Pre = head
	head.Next = node
}
func removeNode(node *Node) {

	// a node c
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}
func removeTail(tail *Node) {
	// a b tail
	tail.Pre.Pre.Next = tail
	tail.Pre = tail.Pre.Pre
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{size: capacity, values: map[int]*Node{}, head: head, tail: tail}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.values[key]; ok {
		removeNode(v)
		addFirst(this.head, v)
		return v.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.values[key]; ok {
		removeNode(v)
		addFirst(this.head, v)
		this.values[key].Key = key
		this.values[key].Val = value
		return
	}
	if this.count == this.size {
		delete(this.values, this.tail.Pre.Key)
		removeTail(this.tail)
		this.count--
	}

	vNode := &Node{Val: value, Key: key}
	this.values[key] = vNode
	addFirst(this.head, vNode)

	this.count++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
