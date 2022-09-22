package unique

import (
	"fmt"
	"testing"
)

type GetUniqueTestCase struct {
	name      string
	value     []int
	expected1 int
	expected2 string
}

func TestGetUnique(t *testing.T) {
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}
	nums4 := []int{1, 1, 1}

	t.Run("test bubble sort", func(t *testing.T) {
		tests := []GetUniqueTestCase{
			{name: "Testing value: [2, 2, 1]", value: nums1, expected1: 1, expected2: ""},
			{name: "Testing value: [4, 1, 2, 1, 2]", value: nums2, expected1: 4, expected2: ""},
			{name: "Testing value: [1]", value: nums3, expected1: 1, expected2: ""},
			{name: "Testing value: [1, 1, 1]", value: nums4, expected1: -1, expected2: "not found"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				theNumber, theText := GetUnique(test.value)

				if theNumber != test.expected1 && theText != test.expected2 {
					fmt.Println("TestGetUnique")
					fmt.Println(test.value)
					fmt.Println(test.expected1)
					t.Fail()
				}
			})
		}
	})
}
