package main

import (
	"code-kata/exercises"
	"fmt"
)

func main() {
	intervals := [][]int{
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

	result := exercises.MergeIntervals(intervals)
	fmt.Println(result)

	result2 := exercises.MergeIntervals(intervals2)
	fmt.Println(result2)

	result3 := exercises.MergeIntervals(intervals3)
	fmt.Println(result3)
}

/*
	input: nums = [2,7,11,15], target = 9
	output: [0,1]
	explanation: because nums[0] + nums[1] == 9, we return [0,1]

	input: nums = [3,2,4], target = 6
	output: [1,2]

	input: nums = [3,3], target = 6
	output: [0,1]
*/

func findIndex(nums []int, target int) {
	var firstIndex int
	var secondIndex int
	// var output string

	for i := 0; i < len(nums); i++ {
		firstIndex = i
		secondIndex = i + 1

		if nums[i] == target {
			fmt.Println(i)
		}

		if secondIndex != len(nums) {
			if nums[firstIndex]+nums[secondIndex] == target {
				// fmt.Println("voila")
				// fmt.Printf("firstIndex: %d\n", firstIndex)
				// fmt.Printf("secondIndex: %d\n", secondIndex)
				fmt.Printf("[%d, %d] \n", firstIndex, secondIndex)
				return
			} else {
				fmt.Println("[]")
			}
		}

	}
}
