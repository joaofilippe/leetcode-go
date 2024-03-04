package reverse

import "math"

func reverse(x int) int {
	reverse := 0
	for x != 0 {
		r := x % 10
		x /= 10
		reverse = reverse*10 + r

		if reverse < math.MinInt32 || reverse >= math.MaxInt32 {
			return 0
		}
	}
	
	return reverse
}
