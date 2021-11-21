package main

type Node struct {
	Val      int
	Children []*Node
}

func main() {

}

func maxDepth(root *Node) (res int) {
	if root == nil {
		return 0
	}
	for _, child := range root.Children {
		if temp := maxDepth(child); res < temp {
			res = temp
		}
	}
	return 1 + res
}
