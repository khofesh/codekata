package exercises

import (
	"fmt"
	"reflect"
	"testing"
)

type MergeIntervalsTestCase struct {
	name     string
	value    [][]int
	expected [][]int
}

func TestMergeIntervals(t *testing.T) {
	var intervals [][]int = [][]int{
		{1, 4},
		{10, 11},
		{4, 5},
		{12, 16},
		{14, 16},
	}
	intervals2 := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	intervals3 := [][]int{
		{1, 4},
		{4, 5},
	}
	expected := [][]int{
		{1, 5},
		{10, 11},
		{12, 16},
	}
	expected2 := [][]int{
		{1, 6},
		{8, 10},
		{15, 18},
	}
	expected3 := [][]int{
		{1, 5},
	}

	t.Run("test merge intervals", func(t *testing.T) {
		tests := []MergeIntervalsTestCase{
			{name: "Testing value: [[1 4] [10 11] [4 5] [12 16] [14 16]]", value: intervals, expected: expected},
			{name: "Testing value: [[1 3] [2 6] [8 10] [15 18]]", value: intervals2, expected: expected2},
			{name: "Testing value: [[1 4] [4 5]]", value: intervals3, expected: expected3},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				mergeResult := MergeIntervals(test.value)

				if !reflect.DeepEqual(test.expected, mergeResult) {
					fmt.Println(test.value)
					fmt.Println(test.expected)
					t.Fail()
				}
			})
		}
	})
}
