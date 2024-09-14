package main

import "fmt"

// 0->1->2->3->4->5

type Node struct {
	value int
	next *Node
}


func createLinkedList(n int) *Node{
	if n <= 0 {
		return nil
	}

	head := &Node{value: 0, next: nil}

	currentNode := head

	for i:=1; i < n; i++ {
		newNode := &Node{value: i, next: nil}
		currentNode.next = newNode
		currentNode = newNode
	}

	return head
}

func printLinkedList(head *Node){
	currentNode := head
	
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.value)
		currentNode = currentNode.next
	}

	fmt.Println("nil")
}
func main(){

	head := createLinkedList(1000000)

	printLinkedList(head)

	// the Go garbage collector will automatically clean up the memory once the program finished
	// or the linkedList is no longer in use

}

