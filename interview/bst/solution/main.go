// 3: Node struct
// 1: BST struct
// 5: BST.Insert
// 17: Node.insert
// 1: inOrder
// 5: inOrderTraversal
// 1: Search
// 13: searchNode

// Mistakes
// First try: --

// Use these variable names: Node = root, newNode, currNode; BST = bst; string = item, key
// Update base package if you changed something here

package main

import (
	"fmt"
	"errors"
)

type Node struct {
	item string
	left *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(item string) {
	if bst.root == nil {
		bst.root = &Node{item, nil, nil}
	} else {
		bst.root.insert(&Node{item, nil, nil})
	}
}

func (currNode *Node) insert(newNode *Node) error {
	if newNode.item == currNode.item {
		return errors.New("Duplicate")
	}

	if newNode.item < currNode.item {
		if currNode.left != nil {
			currNode.left.insert(newNode)
		} else {
			currNode.left = newNode
		}
	} else {
		if currNode.right != nil {
			currNode.right.insert(newNode)
		} else {
			currNode.right = newNode
		}
	}

	return nil
}

func (bst *BST) inOrder() {
	bst.inOrderTraversal(bst.root)
}

func (bst *BST) inOrderTraversal(node *Node) {
	if node != nil {
		bst.inOrderTraversal(node.left)
		fmt.Println(node.item)
		bst.inOrderTraversal(node.right)
	}
}


func (bst *BST) Search (key string) *Node {
	return bst.searchNode(bst.root, key)
}

func (bst *BST) searchNode (currNode *Node, key string) *Node {
	if currNode == nil {
		return nil
	}
	if currNode.item == key {
		return currNode
	}
	if currNode.item > key {
		return bst.searchNode(currNode.left, key)
	} else {
		return bst.searchNode(currNode.right, key)
	}
}

func (bst *BST) preOrderTraverse(t *Node) {
	if t != nil {
		fmt.Println(t.item)
		bst.preOrderTraverse(t.left)
		bst.preOrderTraverse(t.right)
	}
}
func (bst *BST) preOrder() {
	bst.preOrderTraverse(bst.root)
}
func (bst *BST) postOrderTraverse(t *Node) {
	if t != nil {
		bst.postOrderTraverse(t.left)
		bst.postOrderTraverse(t.right)
		fmt.Println(t.item)
	}
}
func (bst *BST) postOrder() {
	bst.postOrderTraverse(bst.root)
}

func main() {
	// Section 1
	bst := &BST{nil}
	eg := []string{"Andy", "Aiken", "Zander", "Jaina", "Mandy"}
	for _, item := range eg {
		bst.Insert(item)
	}

	// Section 2
	fmt.Println("InOrder Traversal...")
	bst.inOrder()


	// Section 3
	fmt.Println("Testing search...")
	item := "Mandy"
	t := bst.Search(item)
	if t == nil {
		fmt.Printf("Node %+v not found\n", item)
	} else {
		fmt.Printf("Node %+v found.\n", item)
	}
}
