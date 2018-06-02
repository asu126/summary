/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result []int
func postorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
    tree(root)
    return result  
}

func tree(root *TreeNode){
    if root == nil {
        return   
    }
    tree(root.Left)
    tree(root.Right)
    result = append(result, root.Val)
}

