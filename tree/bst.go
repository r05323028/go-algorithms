package tree

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (t *BinarySearchTree) Insert(value int) {
	if t.root == nil {
		t.root = &Node{
			value: value,
		}
	} else {
		t.root.Insert(value)
	}
}

func (t *BinarySearchTree) Min() int {
	return t.root.Min()
}

func (t *BinarySearchTree) Max() int {
	return t.root.Max()
}

func (n *Node) Insert(value int) {
	if n.value > value {
		if n.left == nil {
			n.left = &Node{
				value: value,
			}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{
				value: value,
			}
		} else {
			n.right.Insert(value)
		}
	}
}

func (n *Node) Min() int {
	if n.left == nil {
		return n.value
	}
	return n.left.Min()
}

func (n *Node) Max() int {
	if n.right == nil {
		return n.value
	}
	return n.right.Max()
}
