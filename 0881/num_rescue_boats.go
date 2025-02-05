package numrescueboats

import "github.com/joaofilippe/leetcode-go/algorithms/sorts"

func numRescueBoats(people []int, limit int) int {
	boats := 0

	sortedPeople := sorts.QuickSort(people)

	heavier := len(sortedPeople) - 1
	lighter := 0

	for heavier >= lighter {
		if sortedPeople[heavier]+sortedPeople[lighter] <= limit {
			boats++
			heavier--
			lighter++
		} else {
			boats++
			heavier--
		}
	}

	return boats
}
