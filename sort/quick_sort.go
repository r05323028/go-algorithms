package sort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot, arr := arr[0], arr[1:]
	right := []int{}
	left := []int{}
	for _, i := range arr {
		if i <= pivot {
			left = append(left, i)
		} else {
			right = append(right, i)
		}
	}
	return append(append(QuickSort(left), pivot), QuickSort(right)...)

}
