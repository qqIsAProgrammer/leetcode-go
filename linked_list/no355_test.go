package linked_list

import (
	"fmt"
	"testing"
)

func TestTweet(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	twitter.Follow(1, 2)
	twitter.Follow(2, 1)
	twitter.PostTweet(2, 6)
	fmt.Println(twitter.GetNewsFeed(1))
	fmt.Println(twitter.GetNewsFeed(2))
}
