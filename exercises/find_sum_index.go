package exercises

/*
	input: nums = [2,7,11,15], target = 9
	output: [0,1]
	explanation: because nums[0] + nums[1] == 9, we return [0,1]

	input: nums = [3,2,4], target = 6
	output: [1,2]

	input: nums = [3,3], target = 6
	output: [0,1]
*/
func FindSumIndex(nums []int, target int) (int, int) {
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
