package main

import "fmt"

//TreeNode structure defined for Tree Node
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//BinarySearchTree defines the Binary Search Tree
type BinarySearchTree struct {
	root *TreeNode
}

//InsertElement method to insert a node in Binary Search Tree
func (tree *BinarySearchTree) InsertElement(value int) {
	var treeNode *TreeNode
	treeNode = &TreeNode{value, nil, nil}
	if tree.root == nil {
		tree.root = treeNode
	} else {
		insertNodeUtil(tree.root, treeNode)
	}
}

func insertNodeUtil(root *TreeNode, node *TreeNode) {
	if node.val < root.val {
		if root.left == nil {
			root.left = node
		} else {
			insertNodeUtil(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insertNodeUtil(root.right, node)
		}
	}
}

//InOrderTraversal method to do inorder traversal of tree
func (tree *BinarySearchTree) InOrderTraversal() {
	inOrderTraverseUtil(tree.root)
}

func inOrderTraverseUtil(root *TreeNode) {
	if root != nil {
		inOrderTraverseUtil(root.left)
		fmt.Printf("%d ", root.val)
		inOrderTraverseUtil(root.right)
	}
}

func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(4)
	tree.InsertElement(2)
	tree.InsertElement(1)
	tree.InsertElement(3)
	tree.InsertElement(6)
	tree.InsertElement(5)

	tree.InOrderTraversal()
	fmt.Printf("\n")
}
