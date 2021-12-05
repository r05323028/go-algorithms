package sorts

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i
		for ; j > 0 && arr[j-1] >= temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
	return arr
}
