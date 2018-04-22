/**
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMin(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func minDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    
    if root.Left ==nil {
       return minDepth(root.Right) + 1
    }
                     
    if root.Right == nil{
       return minDepth(root.Left) + 1
    }

    return getMin(minDepth(root.Left),minDepth(root.Right)) + 1
}
