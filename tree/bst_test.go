package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := BinarySearchTree{
		root: &Node{value: 80},
	}

	equal := reflect.DeepEqual(tree.root, &Node{value: 80})
	if !equal {
		t.Errorf("tree.root != &Node{value: 80}")
	}
}
func TestInsert(t *testing.T) {
	tree := BinarySearchTree{
		root: &Node{value: 90},
	}

	tree.Insert(80)
	tree.Insert(100)

	if tree.root.value != 90 {
		t.Errorf("root != 90")
	}

	if tree.root.left.value != 80 {
		t.Errorf("root.left != 80")
	}

	if tree.root.right.value != 100 {
		fmt.Println(tree.root.right.value)
		t.Errorf("root.right != 100")
	}
}

func TestMin(t *testing.T) {
	tree := BinarySearchTree{
		root: &Node{value: 90},
	}
	tree.Insert(80)
	tree.Insert(100)

	if tree.Min() != 80 {
		t.Errorf("tree.Min() != 80")
	}
}

func TestMax(t *testing.T) {
	tree := BinarySearchTree{
		root: &Node{value: 90},
	}
	tree.Insert(80)
	tree.Insert(100)

	if tree.Max() != 100 {
		t.Errorf("tree.Max() != 100")
	}
}

func TestNewNode(t *testing.T) {
	n := &Node{value: 80}

	equal := reflect.DeepEqual(n, &Node{value: 80})
	if !equal {
		t.Errorf("n != Node{value: 80}")
	}
}

func TestNodeInsert(t *testing.T) {
	n := Node{value: 90}

	n.Insert(80)
	n.Insert(100)

	if n.value != 90 {
		t.Errorf("node.value != 90")
	}

	if n.left.value != 80 {
		t.Errorf("node.left.value != 80")
	}

	if n.right.value != 100 {
		t.Errorf("node.right.value != 100")
	}
}

func TestNodeMin(t *testing.T) {
	n := Node{value: 90}

	n.Insert(80)
	n.Insert(100)

	if n.Min() != 80 {
		t.Errorf("node.Min() != 80")
	}
}

func TestNodeMax(t *testing.T) {
	n := Node{value: 90}

	n.Insert(80)
	n.Insert(100)

	if n.Max() != 100 {
		t.Errorf("node.Max() != 100")
	}
}
