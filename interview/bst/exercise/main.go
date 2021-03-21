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
	if currNode.item == newNode.item {
		return errors.New("Duplicate found")
	}

	if currNode.item > newNode.item {
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

func (bst *BST) InOrder() {
	bst.inOrder(bst.root)
}

func (bst *BST) inOrder(currentNode *Node) {
	if currentNode != nil {
		bst.inOrder(currentNode.left)
		fmt.Println(currentNode.item)
		bst.inOrder(currentNode.right)
	}
}

func (bst *BST) Search(item string) *Node {
	return bst.search(bst.root, item)
}

func (bst *BST) search(currNode *Node, item string) *Node {
	if currNode == nil {
		return nil
	}
	if currNode.item == item {
		return currNode
	}
	if currNode.item > item {
		return bst.search(currNode.left, item)
	} else {
		return bst.search(currNode.right, item)
	}
}

func main() {
	bst := &BST{nil}
	eg := []string{"Andy", "Aiken", "Zander", "Jaina", "Mandy"}
	for _, item := range eg {
		bst.Insert(item)
	}

	// Section 2
	fmt.Println("InOrder Traversal...")
	bst.InOrder()

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