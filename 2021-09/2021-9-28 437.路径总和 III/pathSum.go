package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

func dfs(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var res int
	if targetSum == root.Val {
		res++
	}

	res += dfs(root.Left, targetSum-root.Val)
	res += dfs(root.Right, targetSum-root.Val)
	return res
}
