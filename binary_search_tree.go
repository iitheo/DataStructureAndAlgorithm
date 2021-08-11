package main

import "fmt"

type Node struct{
	Key int
	Left *Node
	Right *Node
}
var count int

func main(){
	tree := &Node{Key:100}
	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)
	fmt.Println(tree)

	fmt.Println(tree.Search(76))
	fmt.Println(count)
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}
	if n.Key < k {
		//move right
		return n.Right.Search(k)
	}else if n.Key > k {
		//move left
		return n.Left.Search(k)
	}
	return true
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key:k}
		}else{
			n.Right.Insert(k)
		}
	}else if n.Key > k {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: k}
		}else {
			n.Left.Insert(k)
		}
	}
}
