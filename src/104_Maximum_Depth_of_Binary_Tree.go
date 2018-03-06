/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left_hight := maxDepth(root.Left)
    rigth_hight := maxDepth(root.Right)
    if left_hight >= rigth_hight {
        return left_hight + 1
    }else{
        return rigth_hight + 1
    }
}
