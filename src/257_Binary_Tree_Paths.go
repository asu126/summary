package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var count int
var restlt []string

func binaryTreePaths(root *TreeNode) []string {
	count = 0
	restlt = make([]string, 0, 100)
	dfs(root, "")
	return restlt
}
func dfs(root *TreeNode, s string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		s += strconv.Itoa(root.Val)
		restlt = restlt[:count+1]
		restlt[count] = s
		count++
	}
	// fmt.Println(root.Val)
	s += strconv.Itoa(root.Val) + "->"
	dfs(root.Left, s)
	dfs(root.Right, s)
}

func main() {
	// array := []int{1, 0, 0, 0, 1}

	l5 := &TreeNode{
		5,
		nil,
		nil,
	}

	l2 := &TreeNode{
		2,
		nil,
		l5,
	}
	l3 := &TreeNode{
		3,
		nil,
		nil,
	}
	l1 := &TreeNode{
		1,
		l2,
		l3,
	}

	array := make([]string, 0, 20)
	array = array[:2]
	array[0] = "sss"
	array[1] = "111"
	fmt.Println(array)
	fmt.Println(len(array))
	var s []string
	copy(s, array)
	fmt.Println(s)

	aa := binaryTreePaths(l1)
	fmt.Println(aa)
	fmt.Println(strconv.Itoa(2))
}
