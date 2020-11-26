// 3: Node struct
// 1: BST struct
// 5: BST.Insert
// 17: Node.insert
// 1: InOrder
// 5: inOrder
// 1: Search
// 13: searchNode

// Mistakes
// 1x Argument of BST insert should be *Node (easier to expand node object if modification is needed)
// 1x Returning errors are not necessary for first insert
// 1x Did not account for same item string when inserting
// 1x inOrder() should contain node as argument
// 1x Both inOrder and InOrder should have *BST as parent
// 1x bst.Search should return a Node
// 1x node.search should have twoo methods
// 1x a function with return object should contain return within the function

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

func (bst *BST) inOrder(currNode *Node) {
	if currNode != nil {
		bst.inOrder(currNode.left)
		fmt.Println(currNode.item)
		bst.inOrder(currNode.right)
	}
}


func (bst *BST) Search (item string) *Node {
	return bst.searchNode(bst.root, item)
}

func (bst *BST) searchNode (currNode *Node, item string) *Node {
	if currNode == nil {
		return nil
	}
	if currNode.item == item {
		return currNode
	}
	if currNode.item > item {
		return bst.searchNode(currNode.left, item)
	} else {
		return bst.searchNode(currNode.right, item)
	}
}

func (bst *BST) preOrderTraverse(currNode *Node) {
	if currNode != nil {
		fmt.Println(currNode.item)
		bst.preOrderTraverse(currNode.left)
		bst.preOrderTraverse(currNode.right)
	}
}
func (bst *BST) preOrder() {
	bst.preOrderTraverse(bst.root)
}
func (bst *BST) postOrderTraverse(currNode *Node) {
	if currNode != nil {
		bst.postOrderTraverse(currNode.left)
		bst.postOrderTraverse(currNode.right)
		fmt.Println(currNode.item)
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
