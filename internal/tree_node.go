package internal

type TreeNode struct {
	Key    string
	Value  *Word
	Left   *TreeNode
	Right  *TreeNode
	height int
}

func (n *TreeNode) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}

func findMin(n *TreeNode) (string, *Word) {
	if n.Left == nil {
		return n.Key, n.Value
	}
	return findMin(n.Left)
}

func findMax(n *TreeNode) (string, *Word) {
	if n.Right == nil {
		return n.Key, n.Value
	}
	return findMax(n.Right)
}
func (n *TreeNode) getBalanceFactor() int {
	if n.IsLeaf() {
		return 0
	}

	leftTreeHeight := -1
	rightTreeHeight := -1
	if n.Left != nil {
		leftTreeHeight = n.Left.height
	}
	if n.Right != nil {
		rightTreeHeight = n.Right.height
	}
	return leftTreeHeight - rightTreeHeight
}

func (n *TreeNode) rotateLeft() {
	rightNode := n.Right
	rightNodeLeftChild := rightNode.Left
	rightNodeRightChild := rightNode.Right

	newNode := &TreeNode{
		Key:   n.Key,
		Value: n.Value,
		Left:  n.Left,
		Right: rightNodeLeftChild,
	}

	n.Key = rightNode.Key
	n.Value = rightNode.Value
	n.Left = newNode
	n.Right = rightNodeRightChild
}

func (n *TreeNode) rotateRight() {
	leftNode := n.Left
	leftNodeLeftChild := leftNode.Left
	leftNodeRightChild := leftNode.Right

	newNode := &TreeNode{
		Key:   n.Key,
		Value: n.Value,
		Left:  leftNodeRightChild,
		Right: n.Right,
	}

	n.Key = leftNode.Key
	n.Value = leftNode.Value
	n.Left = leftNodeLeftChild
	n.Right = newNode
}
