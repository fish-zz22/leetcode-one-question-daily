package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isEvenOddTree(root *TreeNode) bool {
	nodes, level := []*TreeNode{root}, 0
	for len(nodes) > 0 {
		length, cur := len(nodes), 0
		if level%2 == 1 {
			cur = 1000001
		}
		for i := 0; i < length; i++ {
			node := nodes[0]
			nodes = nodes[1:]
			if level%2 == node.Val%2 || (level%2 == 0 && cur >= node.Val) || (level%2 != 0 && cur <= node.Val) {
				return false
			}
			cur = node.Val
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		level++
	}
	return true
}
