/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(a int) int{
	if a>=0{
		return a
	}else{
		return -a
	}
}

func findTilt(root *TreeNode) int {
	var sum int = 0
	postorder(root, &sum);
	return sum

}

func postorder(node *TreeNode,res *int) int{
	if node==nil{
		return 0;
	} 
	var leftSum int = postorder(node.Left, res);
	var rightSum int = postorder(node.Right, res);
	*res += abs(leftSum - rightSum);
	return leftSum + rightSum + node.Val;
}