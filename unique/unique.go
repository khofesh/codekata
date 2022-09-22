package unique

import "fmt"

func GetUnique(nums []int) (int, string) {
	if len(nums) == 1 {
		return nums[0], ""
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
			return key, ""
		}
	}

	return -1, "not found"
}
