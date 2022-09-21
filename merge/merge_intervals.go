package merge

import (
	"fmt"
	"sort"
)

func MergeIntervals(nums [][]int) [][]int {
	var intervals [][]int = nums

	fmt.Println("before")
	fmt.Println(intervals)

	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		// edge cases
		if len(intervals[i]) == 0 && len(intervals[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(intervals[i]) == 0 || len(intervals[j]) == 0 {
			return len(intervals[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}

		// both slices len() > 0, so can test this now:
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	var previous []int = intervals[0]

	for i := 1; i < len(intervals); i += 1 {
		if previous[1] >= intervals[i][0] {
			previous = []int{previous[0], max(previous[1], intervals[i][1])}
		} else {
			result = append(result, previous)
			previous = intervals[i]
		}
	}

	result = append(result, previous)
	fmt.Println("after")

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
