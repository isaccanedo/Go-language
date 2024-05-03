package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

// Item represents value stored in binary tree node
type Item int

// Node represents one node of binary tree
type Node struct {
	Value Item
	Left  *Node
	Right *Node
}

// BinaryTree represents root node of binary tree
type BinaryTree struct {
	Root *Node
}

// Insert methods inserts new node into binary tree
func (bt *BinaryTree) Insert(value Item) {
	node := &Node{value, nil, nil}
	if bt.Root == nil {
		bt.Root = node
	} else {
		insertNode(bt.Root, node)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func printTree(node *Node, level int) {
	if node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		printTree(node.Left, level)
		fmt.Printf(format+"%d\n", node.Value)
		printTree(node.Right, level)
	}
}

func encodeBinaryTree(bt BinaryTree) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(bt)
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}

func decodeBinaryTree(encodedTree bytes.Buffer) (BinaryTree, error) {
	var newTree BinaryTree
	decoder := gob.NewDecoder(&encodedTree)
	err := decoder.Decode(&newTree)
	return newTree, err
}

func saveBinaryTree(encodedTree bytes.Buffer, filename string) {
	err := os.WriteFile(filename, encodedTree.Bytes(), 0o644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Stored into file\n")
	}
}

func main() {
	var bt BinaryTree
	bt.Insert(5)
	bt.Insert(3)
	bt.Insert(7)
	bt.Insert(1)
	bt.Insert(4)
	bt.Insert(6)
	bt.Insert(8)
	bt.Insert(9)
	bt.Insert(10)
	bt.Insert(0)

	printTree(bt.Root, 0)

	encodedTree, err := encodeBinaryTree(bt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nBuffer with encoded tree: ", encodedTree.Len(), "\n")

	saveBinaryTree(encodedTree, "tree2.gob")

	decodedTree, err := decodeBinaryTree(encodedTree)
	if err != nil {
		fmt.Println(err)
		return
	}
	printTree(decodedTree.Root, 0)
}
