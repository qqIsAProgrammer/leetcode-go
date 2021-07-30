package math

import "math"

func titleToNumber(columnTitle string) int {
	res := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += int(columnTitle[i]-'A'+1) * int(math.Pow(26, float64(len(columnTitle)-1-i)))
	}
	return res
}
