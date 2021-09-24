package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func main() {

}

func flatten(root *Node) *Node {
	dumpy := new(Node)
	dumpy.Next = root

	for root != nil {
		next := root.Next

		if root.Child != nil {
			child := root.Child
			root.Next = child
			child.Prev = root
			root.Child = nil

			for child.Next != nil {
				child = child.Next
			}
			child.Next = next

			if next != nil {
				next.Prev = child
			}
		}

		root = root.Next
	}

	return dumpy.Next
}
