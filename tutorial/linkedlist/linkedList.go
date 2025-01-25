package main

import (
	"fmt"
)

func (l *LinkedList) append(n *Node) {
	// if head pointer is null set the head to the new node
	if l.head == nil {
		l.head = n
		return
	} else {
		// go throught the list and append data
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = n
	}
	l.length++
}

// print elements in the list
func (l *LinkedList) printList() {
	head := l.head
	for head != nil {
		fmt.Printf("Data: %v\n", head.data)
		head = head.next
	}
	fmt.Println("done!")
}

func (l *LinkedList) delete(value int) {
	toDelete := l.head
	for toDelete.next.data != value {
		toDelete = toDelete.next
		return
	}
	if toDelete.next != nil {
		toDelete.next = toDelete.next.next
		l.length--
		return
	}
}

func main() {
	// arr := []int{1, 2, 3, 4, 5}
	p := &Node{}
	n := &Node{data: 10}
	n2 := &Node{data: 20}
	n3 := &Node{data: 30}
	// for _, v := range arr {
	// 	n.data = v
	// }
	myList := LinkedList{}
	myList.append(p)
	myList.append(n)
	myList.append(n2)
	myList.append(n3)
	myList.delete(0)
	// fmt.Println(myList.head)
	fmt.Println(myList.length)
	myList.printList()
}
