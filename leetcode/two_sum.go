package leetcode

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		source, ok := m[target-v]
		if !ok {
			m[v] = i
		} else {
			return []int{source, i}
		}
	}
	return []int{}
}
