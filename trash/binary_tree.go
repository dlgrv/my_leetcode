package main

import (
	"errors"
)

type Node struct {
	Value int
	Data  string
	Left  *Node
	Right *Node
}

func main() {
	n := &Node{10, "aboba", nil, nil}
	n.Insert(2, "")
	n.Insert(14, "")
	n.Insert(5, "")
	n.Insert(8, "")
}

func (n *Node) Insert(value int, data string) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}
	switch {
	case value == n.Value:
		return nil
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(s int) (string, bool) {

	if n == nil {
		return "", false
	}

	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}
