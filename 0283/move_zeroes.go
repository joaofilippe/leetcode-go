package movezeroes

func moveZeroes(nums []int) {
	zeroIdx := 0
	size := len(nums)

	for nonZeroIdx := 0; nonZeroIdx < size; nonZeroIdx++ {
		if nums[nonZeroIdx] != 0 {
			nums[zeroIdx] = nums[nonZeroIdx]
			zeroIdx++
		}
	}

	for i := zeroIdx; i < size; i++ {
		nums[i] = 0
	}
}
