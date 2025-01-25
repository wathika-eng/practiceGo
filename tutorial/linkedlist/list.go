package main

// node contains data and ptr to next node in the list
type Node struct {
	data int
	next *Node
}

// contains head which references first node
type LinkedList struct {
	head   *Node
	length int
}
