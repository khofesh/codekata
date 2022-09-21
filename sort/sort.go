package sort

/*
 * sort slice of int using bubble sort
 */
func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				var lesser int = nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = lesser
			}
		}
	}
}
