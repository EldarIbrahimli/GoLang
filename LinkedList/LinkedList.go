package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	lenght int
}

func (l *linkedList) prepend(n *node) {

	second := l.head
	l.head = n
	l.head.next = second
	l.lenght++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.lenght != 0 {
		fmt.Println("%d ", toPrint.data)
		toPrint = toPrint.next
		l.lenght--
	}
	fmt.Printf("\n")
}

func main() {

	mylist := linkedList{}
	node1 := &node{data: 57}
	node2 := &node{data: 58}
	node3 := &node{data: 59}
	node4 := &node{data: 60}
	node5 := &node{data: 61}
	node6 := &node{data: 62}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)

	mylist.printListData()

}

/*
// A linked list node
type Node struct {
	data int
	next *Node
}

// Constructor to initialize a new node with data
func NewNode(newData int) *Node {
	return &Node{data: newData, next: nil}
}

// Checks whether key is present in linked list
func searchKey(head *Node, key int) bool {
	// Initialize curr with the head of linked list
	curr := head

	// Iterate over all the nodes
	for curr != nil {
		// If the current node's value is equal to key, return true
		if curr.data == key {
			return true
		}
		// Move to the next node
		curr = curr.next
	}

	// If there is no node with value as key, return false
	return false
}

// Driver code
func main() {
	// Create a hard-coded linked list:
	// 14 -> 21 -> 13 -> 30 -> 10
	head := NewNode(14)
	head.next = NewNode(21)
	head.next.next = NewNode(13)
	head.next.next.next = NewNode(30)
	head.next.next.next.next = NewNode(10)

	// Key to search in the linked list
	key := 14

	if searchKey(head, key) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}



package main
import "fmt"

// Node represents a node in a linked list
type Node struct {
   value_num int
   next      *Node
}

//create main function to execute the program
func main() {
   head := &Node{value_num: 10}
   head.next = &Node{value_num: 20}
   head.next.next = &Node{value_num: 30}

   // Accessing elements from linked list
   fmt.Println("Accessing elements of linked list:")
   node := head
   for node != nil {
      fmt.Println(node.value_num)
      node = node.next
   }
}
*/
