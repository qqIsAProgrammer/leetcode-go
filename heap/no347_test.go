package heap

import (
	"log"
	"testing"
)

func TestTopK(t *testing.T) {
	nums := append([]int{}, 1, 1, 1, 2, 2, 3)
	freq := topKFrequent(nums, 2)
	log.Printf("%v\n", freq)
}
