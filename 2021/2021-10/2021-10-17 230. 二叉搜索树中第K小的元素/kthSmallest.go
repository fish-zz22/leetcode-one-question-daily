package main

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res []int
	var dfs func(tree *TreeNode)
	dfs = func(root *TreeNode) {

		if root == nil {
			return
		}

		dfs(root.Left)
		res = append(res, root.Val)
		dfs(root.Right)
	}

	dfs(root)
	return res[k-1]
}
