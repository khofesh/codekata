package main

import "fmt"

func main() {
	nums := []int{9, 7, 11, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}

	var target int = 9
	var target2 int = 6
	var target3 int = 6

	findIndex(nums, target)
	findIndex(nums2, target2)
	findIndex(nums3, target3)
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
