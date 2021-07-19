package linked_list

import "container/heap"

type TweetHeap []*Tweet

func (h TweetHeap) Len() int {
	return len(h)
}

func (h TweetHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h TweetHeap) Less(i, j int) bool {
	return h[i].Timestamp > h[j].Timestamp
}

func (h *TweetHeap) Push(x interface{}) {
	t := x.(*Tweet)
	*h = append(*h, t)
}

func (h *TweetHeap) Pop() interface{} {
	n := len(*h) - 1
	t := (*h)[n]
	*h = (*h)[:n]
	return t
}

var timestamp = 1

type Twitter struct {
	// userId - tweet
	TweetMap map[int]*Tweet
	// follower - followee
	// map[int]interface{}: just a set, and ignore the value
	FollowMap map[int]map[int]interface{}
}

type Tweet struct {
	Id        int
	Timestamp int
	Next      *Tweet
}

func Constructor() Twitter {
	return Twitter{
		TweetMap:  make(map[int]*Tweet),
		FollowMap: make(map[int]map[int]interface{}),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	t := &Tweet{
		Id:        tweetId,
		Timestamp: timestamp,
		Next:      nil,
	}
	timestamp++

	head := this.TweetMap[userId]
	if head == nil {
		this.TweetMap[userId] = t
	} else {
		t.Next = head
		this.TweetMap[userId] = t
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	tweetLists := make([]*Tweet, 0)
	tweetLists = append(tweetLists, this.TweetMap[userId])
	for k := range this.FollowMap[userId] {
		tweetLists = append(tweetLists, this.TweetMap[k])
	}

	return this.getMergedList(tweetLists)
}

func (this *Twitter) getMergedList(lists []*Tweet) []int {
	h := new(TweetHeap)
	heap.Init(h)
	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	news := make([]int, 0)
	for i := 0; i < 10 && h.Len() > 0; i++ {
		t := heap.Pop(h).(*Tweet)
		news = append(news, t.Id)
		if t.Next != nil {
			heap.Push(h, t.Next)
		}
	}

	return news
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.FollowMap[followerId] == nil {
		this.FollowMap[followerId] = make(map[int]interface{})
	}
	// value can be anything
	this.FollowMap[followerId][followeeId] = 0
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	set := this.FollowMap[followerId]
	if set == nil {
		return
	}
	if _, ok := set[followeeId]; ok {
		delete(this.FollowMap[followerId], followeeId)
	}
}
