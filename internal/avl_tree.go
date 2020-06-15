package internal

import "strings"

type AVLTree struct {
	root *TreeNode
}

func (t *AVLTree) Insert(key string, value *Word) {
	t.root = insert(t.root, key, value)
}

func (t *AVLTree) Search(key string) *TreeNode {
	return search(t.root, key)
}

func (t *AVLTree) FuzzySearch(input string, limit int) []*TreeNode {
	nodes := make([]*TreeNode, 0, limit)
	return fuzzySearch(t.root, input, limit, nodes)
}

func insert(root *TreeNode, key string, value *Word) *TreeNode {
	if root == nil {
		return &TreeNode{
			Key:   key,
			Value: value,
		}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key, value)
		updateHeight(root, true)
		ensureBalance(root)
	} else if key > root.Key {
		root.Right = insert(root.Right, key, value)
		updateHeight(root, false)
		ensureBalance(root)
	}

	return root
}

func search(root *TreeNode, key string) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Key {
		return search(root.Left, key)
	}

	if key > root.Key {
		return search(root.Right, key)
	}

	return root
}

func fuzzySearch(root *TreeNode, input string, limit int, nodes []*TreeNode) []*TreeNode {
	if len(nodes) >= limit {
		return nodes
	}

	if root == nil {
		return []*TreeNode{}
	}

	if strings.HasPrefix(root.Key, input) {
		nodes = append(nodes, root)
		return fuzzySearch(root.Left, input, limit, nodes)
	}

	if input < root.Key {
		return fuzzySearch(root.Left, input, limit, nodes)
	}

	if input > root.Key {
		return fuzzySearch(root.Right, input, limit, nodes)
	}

	return []*TreeNode{root}
}

func updateHeight(node *TreeNode, isLeft bool) {
	leftHeight := -1
	rightHeight := -1

	if node.Left != nil {
		leftHeight = node.Left.height
	}
	if node.Right != nil {
		rightHeight = node.Right.height
	}

	if isLeft {
		leftHeight++
	} else {
		rightHeight++
	}

	if leftHeight > rightHeight {
		node.height = leftHeight
	} else {
		node.height = rightHeight
	}
}

func ensureBalance(node *TreeNode) {
	bf := node.getBalanceFactor()
	if bf >= -1 && bf <= 1 {
		return
	}

	if bf > 1 {
		if node.Left.Left == nil {
			node.Left.rotateLeft()
		}
		node.rotateRight()
	} else {
		if node.Right.Right == nil {
			node.Right.rotateRight()
		}
		node.rotateLeft()
	}
}
