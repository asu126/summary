package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var leaf1 []int
var leaf2 []int

var leaf1 []int
var leaf2 []int

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 = make([]int, 0, 100)
	leaf2 = make([]int, 0, 100)

	// 注意指针数组与数组指正的区别
	// 注意*与谁结合，如p *[5]int，*与数组结合说明是数组的指针；如p [5]*int，*与int结合，说明这个数组都是int类型的指针，是指针数组
	dfs(root1, &leaf1)
	dfs(root2, &leaf2)

	if len(leaf1) == len(leaf2) {
		for i := 0; i < len(leaf1); i++ {
			if leaf1[i] != leaf2[i] {
				return false
			}
		}
	} else {
		return false
	}

	return true

}

func dfs(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	}

	dfs(root.Left, leafs)
	dfs(root.Right, leafs)
}

/*
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 = make([]int, 0, 100)
	leaf2 = make([]int, 0, 100)

	dfs(root1, leaf1)
	dfs(root2, leaf2)

	if len(leaf1) == len(leaf2) {
		for i := 0; i < len(leaf1); i++ {
			if leaf1[i] != leaf2[i] {
				return false
			}
		}
	} else {
		return false
	}

	return true

}

func dfs(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		leafs = append(leafs, root.Val)
	}

	dfs(root.Left, leafs)
	dfs(root.Right, leafs)
}
*/
