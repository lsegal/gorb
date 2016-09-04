package node

type Data string

type Node struct {
	Value Data
	Next  *Node
}

func (n *Node) End() bool {
	return n.Next == nil
}
