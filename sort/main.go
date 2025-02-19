package main

import (
	"cmp"
	"fmt"
)

func main() {
	nums := []int{5, 4, 2, 1, 11, 99, 3, 50}
	fmt.Println(quickSort(nums))
	fmt.Println(mergeSort(nums))
}

func quickSort[E cmp.Ordered, S []E](list []E) []E {
	if len(list) <= 1 {
		return list
	}

	if len(list) == 2 {
		if list[0] > list[1] {
			aux := list[1]
			list[1] = list[0]
			list[0] = aux
		}
		return list
	}

	middle := len(list) / 2
	// get pivot from the middle
	pivot := list[middle]
	left := make(S, 0, len(list)-1)
	right := make(S, 0, len(list)-1)

	for i, v := range list {
		// skips if the value is the pivot
		if i == middle {
			continue
		}

		// appends on left values minor than pivot
		if pivot >= v {
			left = append(left, v)
			continue
		}

		// appends on right values bigger than pivot
		right = append(right, v)
	}
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func mergeSort[S ~[]E, E cmp.Ordered](list S) S {
	if len(list) <= 1 {
		return list
	}

	middle := len(list) / 2

	return merge(mergeSort(list[:middle]), mergeSort(list[middle:]))
}

func merge[S ~[]E, E cmp.Ordered](left, right S) S {
	result := make(S, 0, len(left)+len(right))

	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
			continue
		}
		result = append(result, right[r])
		r++
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}
