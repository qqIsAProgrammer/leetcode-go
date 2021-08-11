package dp

import "testing"

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	step := jump(nums)
	t.Log(step)
}
