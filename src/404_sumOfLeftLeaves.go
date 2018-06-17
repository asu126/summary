/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }

    // if root.Left == nil && root.Right == nil {
    //     return root.Val
    // }

    return sumOfLeftLeaves1(root, 0)
}
func sumOfLeftLeaves1(root *TreeNode, sum int) int {
    if root.Left != nil {
        if root.Left.Left == nil && root.Left.Right == nil {
            sum += root.Left.Val
            // fmt.Println(sum)
        } else {
            sum = sumOfLeftLeaves1(root.Left, sum)
        }
    }
    
    if  root.Right != nil {
        sum = sumOfLeftLeaves1(root.Right, sum)
    }
    
    return sum
}

