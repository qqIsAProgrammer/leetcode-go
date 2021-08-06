package graph

import "testing"

func TestFindRedundantConnection(t *testing.T) {
	edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}
	redu := findRedundantConnection(edges)
	for _, v := range redu {
		t.Log(v)
	}
}
