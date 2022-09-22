package exercises

import (
	"testing"
)

type FindSumIndexTestCase struct {
	name      string
	value     []int
	target    int
	expected1 int
	expected2 int
}

func TestFindSumIndex(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	nums1 := []int{2, 11, 7, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}

	t.Run("test 2 nums", func(t *testing.T) {
		tests := []FindSumIndexTestCase{
			{name: "Testing value: [2, 7, 11, 15]", value: nums, target: 9, expected1: 0, expected2: 1},
			{name: "Testing value: [2, 11, 7, 15]", value: nums1, target: 9, expected1: 0, expected2: 3},
			{name: "Testing value: [3, 2, 4]", value: nums2, target: 6, expected1: 1, expected2: 2},
			{name: "Testing value: [3, 3]", value: nums3, target: 6, expected1: 0, expected2: 1},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				res1, res2 := FindSumIndex(test.value, test.target)
				if test.expected1 != res1 && test.expected2 != res2 {
					t.Fail()
				}
			})
		}
	})
}
