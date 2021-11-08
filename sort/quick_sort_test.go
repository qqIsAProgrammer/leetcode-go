package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := append([]int{}, 5, 6, 1, 11, 8, 3, 23, 1, 33, 29, 41, 2)
	QuickSortParallel(nums)
	fmt.Printf("%v\n", nums)
}

func BenchmarkQuickSort(b *testing.B) {
	data := generateRandSlice(10000)
	for i := 0; i < b.N; i++ {
		QuickSort(data)
	}
}

func BenchmarkBuiltinSort(b *testing.B) {
	data := generateRandSlice(20000)
	for i := 0; i < b.N; i++ {
		sort.Ints(data)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	data := generateRandSlice(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(data)
	}
}

func BenchmarkQuickSortParallel(b *testing.B) {
	data := generateRandSlice(5000)
	for i := 0; i < b.N; i++ {
		QuickSortParallel(data)
	}
}

func generateRandSlice(length int) []int {
	rand.Seed(1)
	data := make([]int, 0)
	for i := 0; i < length; i++ {
		data = append(data, rand.Intn(length))
	}
	return data
}

func TestGenerateRandSlice(t *testing.T) {
	data := generateRandSlice(10)
	fmt.Printf("%v\n", data)
}
