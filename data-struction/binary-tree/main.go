package main

import (
	"container/list"
	"fmt"
)

var (
	root *Node
)

// 二叉树节点
type Node struct {
	value int
	left *Node
	right *Node
}

// 初始化1颗二叉树
func init()  {
	node7 := &Node{7, nil, nil}
	node6 := &Node{6, nil, nil}
	node5 := &Node{5, nil, nil}
	node4 := &Node{4, nil, nil}
	node3 := &Node{3, node6, node7}
	node2 := &Node{2, node4, node5}
	node1 := &Node{1, node2, node3}
	root = node1
}

func main() {
	recursionMiddleOrderTraversal(root)
	fmt.Println()
	middleOrderTraversal(root)
}

// 递归方法 实现前序遍历
func recursionPreOrderTraversal(node *Node)  {
	if node != nil {
		fmt.Print(node.value, " ")
		recursionPreOrderTraversal(node.left)
		recursionPreOrderTraversal(node.right)
	}
}

// 递归方法 实现中序遍历
func recursionMiddleOrderTraversal(node *Node)  {
	if node != nil {
		recursionMiddleOrderTraversal(node.left)
		fmt.Print(node.value, " ")
		recursionMiddleOrderTraversal(node.right)
	}
}

// 递归方法 实现中序遍历
func recursionPostOrderTraversal(node *Node)  {
	if node != nil {
		recursionPostOrderTraversal(node.left)
		recursionPostOrderTraversal(node.right)
		fmt.Print(node.value, " ")
	}
}

// 非递归 实现前序遍历
func preOrderTraversal(node *Node) {
	stack := list.New()
	for node != nil || stack.Len() > 0 {
		for node != nil {
			fmt.Print(node.value, " ")
			stack.PushBack(node)
			node = node.left
		}
		if stack.Len() > 0 {
			s := stack.Back()
			stack.Remove(s)
			node = s.Value.(*Node).right
		}
	}
}

// 非递归 实现中序遍历
func middleOrderTraversal(node *Node) {
	stack := list.New()
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.left
		}
		if stack.Len() > 0 {
			s := stack.Back()
			stack.Remove(s)
			fmt.Print(s.Value.(*Node).value, " ")
			node = s.Value.(*Node).right
		}
	}
}
