package main

import "container/heap"

type Tweet struct {
	id    int
	index int
	next  *Tweet
}

type TweetHeap []*Tweet

func (h TweetHeap) Less(i, j int) bool {
	return h[i].index > h[j].index
}

func (h TweetHeap) Len() int {
	return len(h)
}

func (h TweetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *TweetHeap) Push(x any) {
	*h = append(*h, x.(*Tweet))
}

func (h *TweetHeap) Pop() any {
	n := len(*h) - 1
	last := (*h)[n]
	*h = (*h)[:n]
	return last
}

type Twitter struct {
	Post      map[int]*Tweet
	Followees map[int]map[int]int
	Index     int
}

func Constructor() Twitter {
	return Twitter{
		Post:      map[int]*Tweet{},
		Followees: map[int]map[int]int{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Post[userId] = &Tweet{
		tweetId, this.Index, this.Post[userId],
	}
	this.Index++

}

func (this *Twitter) GetNewsFeed(userId int) []int {
	myFeeds := TweetHeap{}
	heap.Init(&myFeeds)

	if this.Post[userId] != nil {
		heap.Push(&myFeeds, this.Post[userId])
	}
	followees := this.Followees[userId]
	for id := range followees {
		if this.Post[id] != nil {
			heap.Push(&myFeeds, this.Post[id])
		}
	}
	feeds := []int{}
	for len(feeds) < 10 && len(myFeeds) > 0 {
		feed := heap.Pop(&myFeeds).(*Tweet)
		if feed.next != nil {
			heap.Push(&myFeeds, feed.next)
		}
		feeds = append(feeds, feed.id)
	}
	return feeds
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.Followees[followerId] == nil {
		this.Followees[followerId] = map[int]int{}
	}
	this.Followees[followerId][followeeId] = 0
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.Followees[followerId], followeeId)
}

/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);
* obj.Unfollow(followerId,followeeId);
 */
