/**
给定一个二叉树，检查它是否是它自己的镜像（即，围绕它的中心对称）。

例如，这个二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是:

    1
   / \
  2   2
   \   \
   3    3
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isSymmetric1(root.Left,root.Right)
    
}

func isSymmetric1(l *TreeNode, r *TreeNode) bool{
    if l == nil && r == nil{
        return true
    }
    if l == nil || r == nil{
        return false
    }
    
    if l.Val != r.Val{
        return false
    }
    
    return isSymmetric1(l.Left,r.Right) && isSymmetric1(l.Right,r.Left)
}