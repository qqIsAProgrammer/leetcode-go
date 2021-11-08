package math

import (
	"log"
	"testing"
)

func TestMaxProduct(t *testing.T) {
	nums := append([]int{}, -4, -3, -2)
	product := maxProduct(nums)
	log.Println(product)
}
