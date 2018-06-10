/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    var count int = 0
    result := getKthSmallest(root, k, &count)
    if result == nil {
        return 0
    }
    
    return result.Val
}

func getKthSmallest(root *TreeNode, k int, count *int) *TreeNode {
    var result  *TreeNode
    if root == nil{
        return nil
    }

    result =  getKthSmallest(root.Left, k, count)
    if result != nil {
        return result
    }
    
    // visit root node
    (*count)++
    if *count == k {
        return root
    }
    
    result =  getKthSmallest(root.Right, k, count)
     if result != nil {
        return result
    }
    return nil
    
}
