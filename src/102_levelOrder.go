/**
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。

例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res [][]int
func levelOrder(root *TreeNode) [][]int {
    res = make([][]int, 0)
    if  root == nil{
        return res
    }
    levelOrder1(root, 0)
    return res
}

func levelOrder1(root *TreeNode, cil int) {
     if  root == nil{
        return
    }
    if root != nil {
        var tmp []int
        // fmt.Println(root.Val)
        // fmt.Println("-------------",len(res), cil)
        if len(res) < (cil+1) {
            tmp = make([]int,0,128)
            res = append(res,tmp)
        }
        tmp = res[cil]
        tmp =append(tmp,root.Val)
        res[cil] = tmp
        // fmt.Println("-------------",res)
    }
    
    levelOrder1(root.Left,cil+1)
    levelOrder1(root.Right,cil+1)
}