package sort

func merge(left []int, right []int) []int {
	i := 0
	j := 0
	var ret []int
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}
	for ; i < len(left); i++ {
		ret = append(ret, left[i])
	}
	for ; j < len(right); j++ {
		ret = append(ret, right[j])
	}
	return ret
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := arr[:mid]
	right := arr[mid:]
	return merge(MergeSort(left), MergeSort(right))
}
