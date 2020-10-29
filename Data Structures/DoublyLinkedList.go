package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	prev *Node
}

func (n *Node) insertHead(head **Node, newData interface{}) {
	newNode := &Node{}
	newNode.data = newData
	newNode.next = *head
	newNode.prev = nil

	if *head != nil {
		(*head).prev = newNode
	}

	*head = newNode
}

func (n *Node) insertAfter(prevNode *Node, newData interface{}) {
	if prevNode == nil {
		_ = errors.New("previous node cannot be empty")
		return
	}

	newNode := &Node{}
	newNode.data = newData
	newNode.next = prevNode.next
	prevNode.next = newNode
	newNode.prev = prevNode

	if newNode.next != nil {
		newNode.next.prev = newNode
	}
}

func (n *Node) insertBefore(head **Node, nextNode *Node, newData interface{}) {
	if nextNode == nil {
		_ = errors.New("next node cannot be empty")
		return
	}

	newNode := &Node{}
	newNode.data = newData
	newNode.prev = nextNode.prev
	newNode.next = nextNode
	nextNode.prev = newNode

	if newNode.prev != nil {
		newNode.prev.next = newNode
	} else {
		*head = newNode
	}
}

func (n *Node) append(head **Node, newData interface{}) {
	newNode := &Node{}
	lastNode := *head
	newNode.data = newData
	newNode.next = nil

	if *head == nil {
		newNode.prev = nil
		*head = newNode
		return
	}

	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode
	newNode.prev = lastNode
	return
}

func (n *Node) deleteHead(head **Node) {
	if head == nil {
		_ = errors.New("head node cannot be empty")
		return
	}

	delNode := *head
	*head = (*head).next

	if delNode.next != nil {
		delNode.next.prev = nil
	}

	delNode = nil
	return
}

func (n *Node) deleteAfter(prevNode *Node) {
	if prevNode == nil {
		_ = errors.New("previous node cannot be empty")
		return
	}

	delNode := prevNode.next
	prevNode.next = delNode.next

	if delNode.next != nil {
		delNode.next.prev = prevNode
	}

	delNode = nil
	return
}

func (n *Node) deleteNode(head **Node) {
	if head == nil {
		_ = errors.New("list cannot be empty")
		return
	}

	lastNode := *head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.prev.next = nil
	lastNode = nil
	return
}

func (n *Node) deleteValue(head **Node, delValue interface{}) {
	if head == nil {
		_ = errors.New("list cannot be empty")
		return
	}

	delNode := *head

	for delNode != nil {
		if delNode.data == delValue {
			if delNode == *head {
				*head = delNode.next
				delNode.next.prev = nil
				delNode = nil
				return
			}
			delNode.prev.next = delNode.next
			if delNode.next != nil {
				delNode.next.prev = delNode.prev
			}
			delNode = nil
			return
		} else {
			delNode = delNode.next
		}
	}

	_ = errors.New("value not found")
	return
}

func (n *Node) len(head **Node) int {
	node := *head
	length := 0
	for node != nil {
		node = node.next
		length += 1
	}

	return length
}

func (n *Node) isEmpty(head **Node) bool {
	if (*head).len(head) == 0 {
		return true
	} else {
		return false
	}
}

func (n *Node) getIndex(head **Node, value interface{}) int {
	valueNode := *head
	index := 0

	for valueNode != nil {
		if valueNode.data == value {
			return index
		} else {
			index += 1
			valueNode = valueNode.next
		}
	}
	return -1
}

func printList(node *Node) {
	lastNode := &Node{}
	fmt.Printf("\nForward \n")
	for node != nil {
		fmt.Printf("%d \t", node.data)
		lastNode = node
		node = node.next
	}

	fmt.Printf("\nBackward \n")
	for lastNode != nil {
		fmt.Printf("%d \t", lastNode.data)
		lastNode = lastNode.prev
	}
}

func main() {
	var head *Node

	fmt.Printf("\nIs the list empty? %t\n", head.isEmpty(&head))

	head.append(&head, 6)
	printList(head)

	head.insertHead(&head, 7)
	printList(head)

	head.insertHead(&head, 1)
	printList(head)

	head.append(&head, 4)
	printList(head)

	head.insertAfter(head.next, 8)
	printList(head)

	head.insertBefore(&head, head.next.next, 5)
	printList(head)

	head.deleteAfter(head.next)
	printList(head)

	head.deleteHead(&head)
	printList(head)

	head.deleteNode(&head)
	printList(head)

	head.append(&head, 9)
	printList(head)

	head.append(&head, 2)
	printList(head)

	head.deleteValue(&head, 7)
	printList(head)

	value := 8
	fmt.Printf("\nIs the list empty? %t\n", head.isEmpty(&head))
	fmt.Printf("\nThe length of the list %d\n", head.len(&head))
	fmt.Printf("\nThe index of value %d is %d\n", value, head.getIndex(&head, value))
}
