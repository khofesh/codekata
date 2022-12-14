package main

import "fmt"

func main() {
	nums1 := []int{2, 2, 1}
	nums2 := []int{4, 1, 2, 1, 2}
	nums3 := []int{1}

	res1 := getUnique(nums1)
	fmt.Println(res1)
	res2 := getUnique(nums2)
	fmt.Println(res2)
	res3 := getUnique(nums3)
	fmt.Println(res3)

}

func getUnique(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	numToCount := make(map[int]int, len(nums))
	for _, num := range nums {
		if _, ok := numToCount[num]; ok {
			numToCount[num] = 1
		}
		numToCount[num] += 1
	}

	fmt.Println(numToCount)

	for key, element := range numToCount {
		if element == 1 {
			return key
		}
	}

	return -1
}
