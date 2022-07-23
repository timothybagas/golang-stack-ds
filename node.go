package stack

type node struct {
	next *node
	val  interface{}
}

func newNode(next *node, val interface{}) *node {
	return &node{next, val}
}
