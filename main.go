package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	nums1 := []int{2, 11, 7, 15}
	nums2 := []int{3, 2, 4}
	nums3 := []int{3, 3}

	var target int = 9
	var target2 int = 6
	var target3 int = 6

	findIndex(nums, target)
	findIndex(nums2, target2)
	findIndex(nums3, target3)

	index1, index2 := twoSum2(nums1, target)
	fmt.Printf("index1: %d\n", index1)
	fmt.Printf("index2: %d\n", index2)
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
			fmt.Printf("[%d]\n", i)
		}

		if secondIndex != len(nums) {
			if nums[firstIndex]+nums[secondIndex] == target {
				// fmt.Println("voila")
				// fmt.Printf("firstIndex: %d\n", firstIndex)
				// fmt.Printf("secondIndex: %d\n", secondIndex)
				fmt.Printf("[%d, %d] \n", firstIndex, secondIndex)
			}
		}
	}
}

func findSumIndex(nums []int, target int) (int, int) {
	if len(nums) <= 1 {
		return -1, -1
	}

	numToIndex := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numToIndex[num]; ok {
			return j, i
		}
		numToIndex[target-num] = i
	}

	return -1, -1
}

// func getUnique(nums []int) int {
// 	if len(nums) == 1 {
// 		return 0
// 	}

// 	numToIndex := make(map[int]int, len(nums))
// 	for i, num := range nums {

// 	}
// }
