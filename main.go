package main

import (
	remove_duplicate "github.com/joaofilippe/leetcode-go/0026"
)

func main() {
	r := remove_duplicate.RemoveDuplicates([]int{1, 1, 2})
	println(r == 2)

	r = remove_duplicate.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	println(r == 5)
}
