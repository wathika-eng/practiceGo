package main

import "fmt"

func (l *LinkedList) append(n *Node) {
	// if head pointer is null
	if l.head == nil {
		l.head = n
		l.length++
	}
	l.head.next = n
	l.length++
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := &Node{}
	// n2 := &Node{data: 20}
	for _, v := range arr {
		n.data = v
	}
	myList := LinkedList{}
	myList.append(n)
	// myList.append(n2)
	fmt.Println(myList.head)
	fmt.Println(myList.length)
}
