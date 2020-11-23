// TODO count nodes and remove nodes

package main

import "fmt"

// lines of code : name
// 3: BinaryNode
// 1: BST

// main "Andy", "Aiken", "Zander", "Jaina", "Mandy"

type BinaryNode struct {
	item  string
	left  *BinaryNode
	right *BinaryNode
}

type BST struct {
	root *BinaryNode
}

func (root *BinaryNode) insert(new_node *BinaryNode) {
	if new_node.item > root.item {
		if root.right == nil {
			root.right = new_node
		} else {
			root.right.insert(new_node)
		}
	} else if new_node.item < root.item {
		if root.left == nil {
			root.left = new_node
		} else {
			root.left.insert(new_node)
		}
	}
}

func (bst *BST) Insert(item string) {
	if bst.root == nil {
		bst.root = &BinaryNode{item, nil, nil}
	}
	bst.root.insert(&BinaryNode{item, nil, nil})
}

func (bst *BST) inOrderTraverse(t *BinaryNode) {
	if t != nil {
		bst.inOrderTraverse(t.left)
		fmt.Println(t.item)
		bst.inOrderTraverse(t.right)
	}
}

func (bst *BST) inOrder() {
	bst.inOrderTraverse(bst.root)
}

func (bst *BST) searchNode(t *BinaryNode, item string) *BinaryNode {
	if t == nil {
		return nil
	} else {
		if t.item == item {
			return t
		} else {
			if item < t.item {
				return bst.searchNode(t.left, item)
			} else {
				return bst.searchNode(t.right, item)
			}
		}
	}
}
func (bst *BST) search(item string) *BinaryNode {
	return bst.searchNode(bst.root, item)
}


func main() {
	bst := &BST{nil}
	eg := []string{"Andy", "Aiken", "Zander", "Jaina", "Mandy"}
	for _, item := range eg {
		bst.Insert(item)
	}
	fmt.Println("InOrder Traversal...")
	bst.inOrder()

	fmt.Println("Testing search...")
	item := "Mandy"
	t := bst.search(item)
	if t == nil {
		fmt.Printf("Node %+v not found\n", item)
	} else {
		fmt.Printf("Node %+v found.\n", item)
	}
}
