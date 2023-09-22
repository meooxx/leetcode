package main

import "container/heap"

type Tweet struct {
	Id    int
	Index int
	Next  *Tweet
}

type TweetHeap []*Tweet

func (this TweetHeap) Len() int {
	return len(this)
}
func (this TweetHeap) Less(a, b int) bool {
	return this[a].Index > this[b].Index
}
func (this TweetHeap) Swap(a, b int) {
	this[a], this[b] = this[b], this[a]
}

func (this *TweetHeap) Push(x any) {
	*this = append(*this, x.(*Tweet))
}

func (this *TweetHeap) Pop() any {
	l := len(*this)
	last := (*this)[l-1]
	*this = (*this)[:l-1]
	return last
}

type Twitter struct {
	Post     map[int]*Tweet
	Followee map[int]map[int]int
	count    int
}

func Constructor() Twitter {

	return Twitter{
		Post:     map[int]*Tweet{},
		Followee: map[int]map[int]int{},
		count:    0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Post[userId] = &Tweet{
		Id:    tweetId,
		Next:  this.Post[userId],
		Index: this.count,
	}
	this.count++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	heap.Init(h)
	if this.Post[userId] != nil {
		heap.Push(h, this.Post[userId])
	}
	for v := range this.Followee[userId] {
		if this.Post[v] != nil {
			heap.Push(h, this.Post[v])
		}
	}
	result := []int{}
	for len(result) < 10 && len(*h) > 0 {
		t := heap.Pop(h).(*Tweet)
		if t.Next != nil {
			heap.Push(h, t.Next)
		}
		result = append(result, t.Id)
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.Followee[followerId] == nil {
		this.Followee[followerId] = map[int]int{}
	}
	this.Followee[followerId][followeeId] = 1
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.Followee[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
