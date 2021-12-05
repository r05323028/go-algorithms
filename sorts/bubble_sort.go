package sorts

func BubbleSort(arr []int) []int {
	swap := true
	for swap {
		swap = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swap = true
			}
		}
	}
	return arr
}
