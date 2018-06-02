/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result []int
func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
    tree(root)
    return result  
}

func tree(root *TreeNode){
    if root == nil {
        return   
    }
    result = append(result, root.Val)
    tree(root.Left)
    tree(root.Right)
}


// zhizhneg
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	var result []int = make([]int, 0)
    tree(root, &result)
    return result  
}

func tree(root *TreeNode, result *[]int){
    if root == nil {
        return   
    }
    tree(root.Left, result)
    tree(root.Right, result)
    *result = append(*result, root.Val)
}
