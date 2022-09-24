package main

import (
	"fmt"
	"bytes"
	"errors"
)

/**
 * s string "abc"
 * 
 * if('0' < s[i] < '9') help(s, index + 1) return 
 * else 
 * 	 help("abc", index + 1, r)
 *   help("Abc", index + 1, r)
 * if index == len(s) 
 *   r.push(s)
 * 
 */


 /**
	* bfs 
		 str = abc
		 que = [abc]
     for i<len(str) {
			  len(que) >0
			  for < len(que) {       // i = 0           i = 1     ...
					curr = que.deque()   //  []             [abc]
					curr[i].ToUpper.push // [Abc]           [abc, ABc]
					curr[i].ToLower.push // [Abc, abc]      [abc, ABc, Abc]
				} 									
		 } 
	*
	*
	*    abc
			 Abc 
	*

  */

func main() {
	r := letterCasePermutation("ab")
	fmt.Println(r)

	return 
	q:=&QNode{}
	q.Enque("ab")
	q.Deque()
	q.Enque("Ab")
	q.Enque("ab")
	fmt.Println(q.Front, q.Rear)
	q.Deque()
	q.Enque("AB")
	q.Enque("Ab")
	fmt.Println(q.Front.Next, q.Rear)
}



//DFS
func letterCasePermutation(S string ) []string  {
	res := []string{ }
	//temp := map[string]bool{ }
	// 研究下 backtracking 算法 
	//helper(S, temp, 0)
	v := helperBFS(S)
	for v != nil {
		res = append( res, v.Val )
		v = v.Next
	}
	
	return res
	
}

type Node struct {
	Val string
	Next *Node
}

type QNode struct {
	Front *Node
	Rear *Node
	Size int
}

func(q *QNode) Enque(s string)  {
	node := &Node{ Val:s }
	if q.Front == nil {
		q.Front = node
	}
	q.Size += 1
	if q.Rear == nil {
		q.Rear = node
		return 
	}

	q.Rear.Next = node
	q.Rear = q.Rear.Next

}

func (q *QNode) Deque() (*Node, error) {

	if q.Front == nil {
		return nil, errors.New("empty queue")
	}

	curr := q.Front
	q.Front = q.Front.Next
	q.Size -= 1
	return curr, nil

}




func helperBFS(s string) *Node{
	node := &QNode{}
	node.Enque(s)
	for i:=0;i<len(s);i++ {
		if s[i] >= '0' &&  s[i] <= '9' {
			continue
		}
		size := node.Size		
		for j:=0;j<size;j++ {
			curr, _ := node.Deque()
			node.Enque(UpperAt(curr.Val, i))
			node.Enque(LowerAt(curr.Val, i))
		}
	}
	return node.Front
} 

func helper(s string, temp map[string]bool, index int) {
	
	if index == len(s) {
		temp[s] = true
		return
	}

	if s[index] >='0' && s[index] <= '9' {
		helper(s, temp, index + 1)
		return
	}
	
	
	// upperCase 
	helper(UpperAt(s, index), temp, index + 1)

	//lowerCase 
	helper(LowerAt(s, index), temp, index + 1)
}


func UpperAt(s string, i int) string {
	strs := []byte(s)
	strs[i] = bytes.ToUpper(strs[i:i+1])[0]
	return string(strs)
}
 
func LowerAt(s string, i int) string {
	strs := []byte(s)
	strs[i] = bytes.ToLower(strs[i:i+1])[0]
	return string(strs)
}