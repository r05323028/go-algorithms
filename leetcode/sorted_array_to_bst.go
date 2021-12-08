package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	queue := [](*TreeNode){node}
	ans := []int{}
	for len(queue) > 0 {
		curLevel := []int{}
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[i]
			curLevel = append(curLevel, cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		queue = queue[sz:]
		ans = append(ans, curLevel...)
	}
	return ans
}

func SortedArrayToBST(nums []int) *TreeNode {
	root := &TreeNode{}
	buildTree(nums, root)
	return root
}

func buildTree(nums []int, t *TreeNode) {
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid+1:]
	t.Val = nums[mid]
	if len(left) > 0 {
		t.Left = &TreeNode{}
		buildTree(left, t.Left)
	}
	if len(right) > 0 {
		t.Right = &TreeNode{}
		buildTree(right, t.Right)
	}
}
