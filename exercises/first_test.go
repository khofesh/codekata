package exercises

import (
	"fmt"
	"testing"
)

type TestCase struct {
	name     string
	value    []int
	expected string
	check    int
}

func TestCheckInArray(t *testing.T) {
	nums := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}

	t.Run("test 2 nums", func(t *testing.T) {
		tests := []TestCase{
			{name: "Testing value: [2, 2, 1]", value: nums, expected: "true", check: 1},
			{name: "Testing value: [4, 1, 2, 1, 2]", value: nums2, expected: "false", check: 5},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := CheckInArray(test.value, test.check)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})
}

func TestBinarySearch(t *testing.T) {
	nums := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}

	t.Run("test 2 nums", func(t *testing.T) {
		tests := []TestCase{
			{name: "Testing value: [2, 2, 1]", value: nums, expected: "true", check: 1},
			{name: "Testing value: [4, 1, 2, 1, 2]", value: nums2, expected: "false", check: 5},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := BinarySearch(test.value, test.check)
				if test.expected != actual {
					fmt.Println(actual)

					t.Fail()
				}
			})
		}
	})
}
