package remove_duplicate

func RemoveDuplicates(nums []int) int {
    if len(nums) < 1 {
		return 0
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
            nums[j] = nums[i]
		}
	}

	return j+1
}