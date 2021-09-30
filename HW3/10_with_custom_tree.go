package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func Add(root *TreeNode, value int) (*TreeNode) {
	if (root == nil) {
		node := &TreeNode{Value: value}
		return node
	}
	if (root.Value > value) {
		root.Left = Add(root.Left, value)
	} else if (root.Value < value) {
		root.Right = Add(root.Right, value)
	}
	return root
}

func Inorder(root *TreeNode) {
	if (root == nil) {
		return
	} 
	Inorder(root.Left)
	fmt.Printf("%d ", root.Value)
	Inorder(root.Right)
}

func NewNode(multiplier int) *TreeNode {
	a := []int{}
	for i := 1; i <= 10; i++ {
		a = append(a, i * multiplier)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	var root *TreeNode
	for _, i := range a {
		root = Add(root, i)
	}
	return root
}

func Walk(t *TreeNode, ch chan int) {
    WalkInorder(t, ch)
    close(ch)
}

func WalkInorder(t *TreeNode, ch chan int) {
    if t == nil {
		return
	}
	WalkInorder(t.Left, ch)
	ch <- t.Value
	WalkInorder(t.Right, ch)
}

// Same determines whether the s
// t1 and t2 contain the same values.
func Same(t1, t2 *TreeNode) bool {
    ch1, ch2 := make(chan int), make(chan int)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for {
        n1, ok1 := <- ch1
        n2, ok2 := <- ch2
        if (ok1 != ok2) || (n1 != n2) {
        	return false
        }
        if !ok1 {
        	break
        }
    }
    return true
}

func main() {
    fmt.Println(Same(NewNode(1), NewNode(1)))
    fmt.Println(Same(NewNode(1), NewNode(2)))
}