package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 7, 9, 11, 15, 18, 20}
	fmt.Println(binarySearch(nums, 15))
}

func binarySearch(list []int, value int) int {
	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (left + right) / 2

		if list[mid] == value {
			return mid
		}

		// value is on left side
		if list[mid] > value {
			right = mid - 1
			continue
		}

		// value is on right side
		left = mid + 1
	}

	return -1
}
