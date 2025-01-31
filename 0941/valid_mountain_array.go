package validmountainarray

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 1

	for i < len(arr) && arr[i] > arr[i-1] {
		i++
	}

	if i == 1 || i == len(arr) {
		return false
	}

	for i < len(arr) && arr[i] < arr[i-1] {
		i++
	}

	return i == len(arr)
}
