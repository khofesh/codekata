package pair

import (
	"fmt"
	"reflect"
	"testing"
)

type ZeroLastTestCase struct {
	name     string
	value    []int
	expected []int
}

func TestZeroLast(t *testing.T) {
	num1 := []int{1, 2, 5, 6, 7, 0, 0, 54}
	num2 := []int{1, 2, 0, 0, 0, 3, 6}
	expected1 := []int{1, 2, 5, 6, 7, 54, 0, 0}
	expected2 := []int{1, 2, 3, 6, 0, 0, 0}

	t.Run("test bubble sort", func(t *testing.T) {
		tests := []ZeroLastTestCase{
			{name: "Testing value: [1, 2, 5, 6, 7, 0, 0, 54]", value: num1, expected: expected1},
			{name: "Testing value: [1, 2, 0, 0, 0, 3, 6]", value: num2, expected: expected2},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := ZeroLast(test.value)

				if !reflect.DeepEqual(test.expected, result) {
					fmt.Println("TestZeroLast")
					fmt.Println(result)
					fmt.Println(test.expected)
					t.Fail()
				}
			})
		}
	})
}
