package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2, 5, 8, 10, 11}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		index1 := (low + high) / 2
		index2 := (m+n+1)/2 - index1

		maxLeft1 := math.MinInt64
		if index1 > 0 {
			maxLeft1 = nums1[index1-1]
		}

		minRight1 := math.MaxInt64
		if index1 < m {
			minRight1 = nums1[index1]
		}

		maxLeft2 := math.MinInt64
		if index2 > 0 {
			maxLeft2 = nums2[index2-1]
		}

		minRight2 := math.MaxInt64
		if index2 < n {
			minRight2 = nums2[index2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				return (float64(max(maxRight1, maxRight2)) + float64(min(minLeft1, minLeft2))) / 2.0
			}
			return float64(max(maxRight1, maxRight2))
		} else if maxRight1 > minLeft2 {
			high = index1 - 1
		} else {
			low = index1 + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
