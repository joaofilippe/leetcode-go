package numrescueboats

func quickSort(slc []int) []int {
	sorted := make([]int, len(slc))

	copy(sorted, slc)

	sort(sorted, 0, len(slc)-1)

	return sorted
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIdx := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIdx]

			arr[splitIdx] = arr[i]
			arr[i] = temp

			splitIdx++
		}
	}

	arr[end] = arr[splitIdx]
	arr[splitIdx] = pivot

	sort(arr, start, splitIdx-1)
	sort(arr, splitIdx+1, end)
}
