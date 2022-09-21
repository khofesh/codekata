package exercises

import (
	"fmt"
	"reflect"
	"testing"
)

type BubbleSortTestCase struct {
	name     string
	value    []int
	expected []int
}

func TestBubbleSort(t *testing.T) {
	t.Run("test bubble sort", func(t *testing.T) {
		tests := []BubbleSortTestCase{
			{name: "Testing value: [2, 2, 1]", value: []int{2, 2, 1}, expected: []int{1, 2, 2}},
			{name: "Testing value: [4, 1, 2, 1, 2]", value: []int{4, 1, 2, 1, 2}, expected: []int{1, 1, 2, 2, 4}},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				BubbleSort(test.value)

				if !reflect.DeepEqual(test.expected, test.value) {
					fmt.Println("TestBubbleSort")
					fmt.Println(test.value)
					fmt.Println(test.expected)
					t.Fail()
				}
			})
		}
	})
}
