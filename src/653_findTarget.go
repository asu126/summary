/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var array []int
func findTarget(root *TreeNode, k int) bool {
    array =make([]int, 0)
    dfs(root)
   // fmt.Println(array)

    for i,j := 0, len(array)-1; i<j; {
    	tmp := array[i]+array[j]
    	if tmp == k {
    		return true
    	} else if tmp > k{
    		j--
    	} else {
			i++
    	}
    }
    
    return false
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