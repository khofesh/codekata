package exercises

import (
	sort "codekata/sort"
)

func CheckInArray(nums []int, check int) string {
	var output string = ""

	for _, v := range nums {
		if v == check {
			output = "true"
		}
	}

	if output != "" {
		return output
	} else {
		return "false"
	}
}

func BinarySearch(nums []int, target int) string {
	sort.BubbleSort(nums)
	first := 0
	last := len(nums) - 1

	for first <= last {
		middle := (first + last) / 2

		if nums[middle] < target {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}

	if first == len(nums) || nums[first] != target {
		return "false"
	}

	return "true"
}
