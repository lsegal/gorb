package node

type Data string

type Node struct {
	Value Data
	Next  *Node
}

func New(v Data, n *Node) *Node {
	return &Node{Value: v, Next: n}
}

func (n *Node) End() bool {
	return n.Next == nil
}
