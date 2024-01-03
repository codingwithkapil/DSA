package main

import "fmt"

type Node struct {
	value     int
	left_val  *Node
	right_val *Node
}

func main() {
	root := &Node{value: 10}
	root.left_val = &Node{value: 20}
	root.right_val = &Node{value: 30}
	root.left_val.left_val = &Node{value: 40}
	root.left_val.right_val = &Node{value: 50}

	fmt.Println("The pre-order traversal is given as:")
	preOrder(root)
	inOrder(root)
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left_val)
	fmt.Println(node.value)
	inOrder(node.right_val)
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.value)
	preOrder(node.left_val)
	preOrder(node.right_val)
}
