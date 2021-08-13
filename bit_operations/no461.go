package bit_operations

func hammingDistance(x int, y int) int {
	cnt := 0
	for s := x ^ y; s > 0; s >>= 1 {
		cnt += s & 1
	}
	return cnt
}
