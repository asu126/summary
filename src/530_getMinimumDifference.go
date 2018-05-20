/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var array []int

func getMinimumDifference(root *TreeNode) int {
    array = make([]int, 0)
    dfs(root)

	var length int = len(array)
	if length<2 {
		return 0
	}
	var result int = array[1] - array[0]
    for i := 2; i<length; i++ {
    	tmp := array[i]-array[i-1]
    	if tmp < result {
    		result = tmp
    	}
    }
    
    return result
}

func dfs(root *TreeNode){
	if root == nil{
		return
	}
	if root.Left != nil {
		dfs(root.Left)
	}
	array = append(array, root.Val)
	if root.Right != nil {
		dfs(root.Right)
	}
}
