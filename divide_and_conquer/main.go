package main

import (
	"cmp"
	"fmt"
)

func main() {
	nums := []int{1, 11, 2, 99, 3, 50}
	fmt.Println(sum(nums))
	fmt.Println(count(nums))
	fmt.Println(max(nums))
}

func sum(list []int) int {
	// base case
	if len(list) == 1 {
		return list[0]
	}

	// divide
	return list[0] + sum(list[1:])
}

func count[E any, S []E](list []E) int {
	// base case
	if len(list) == 1 {
		return 1
	}

	// divide
	return 1 + count(list[1:])
}

func max[E cmp.Ordered, S []E](list []E) E {
	// base case
	if len(list) == 2 {
		if list[0] > list[1] {
			return list[0]
		}
		return list[1]
	}

	maxRight := max(list[1:])
	if list[0] > maxRight {
		return list[0]
	}
	return maxRight
}
