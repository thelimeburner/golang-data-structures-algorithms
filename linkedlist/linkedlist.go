package linkedlist

import (
	"errors"
	"fmt"
)

/*
 Sample Linked List Implementation for integers


*/

//Node represents a single node in a list
type Node struct {
	Next *Node
	Data int
}

//LList represents the list of nodes
type LList struct {
	Next *Node
}

//Append adds a node to the list
func (l *LList) Append(value int) {

	//check if first node, if so we need to add the node now.
	if l.Next == nil {
		//create new node
		newNode := Node{Next: nil, Data: value}

		//store node in list
		l.Next = &newNode
		return
	}
	currentPointer := l.Next

	//iterate until we reach end of list
	for currentPointer.Next != nil {
		currentPointer = currentPointer.Next
	}

	//create new node
	newNode := Node{Next: nil, Data: value}

	//store node in list
	currentPointer.Next = &newNode

}

//Del removes an the first occurence of an item from the list by value
func (l *LList) Del(val int) error {
	currentPointer := l.Next

	//if empty list return error
	if currentPointer == nil {
		return errors.New("Empty List, no value to delete")
	}

	//if removing the first node
	if currentPointer.Data == val {
		l.Next = currentPointer.Next
		return nil
	}

	//otherwise track previous and next pointers for a node
	var prevPtr *Node
	var nextPtr *Node

	for currentPointer.Next != nil {
		prevPtr = currentPointer
		currentPointer = currentPointer.Next
		nextPtr = currentPointer.Next

		if currentPointer.Data == val {
			prevPtr.Next = nextPtr
			return nil
		}
	}
	errMsg := fmt.Sprintf("Error: %d not found in List.", val)
	return errors.New(errMsg)
}

//Get returns the value at i
func (l *LList) Get(index int) (int, error) {

	currentPtr := l.Next
	indexPtr := 0
	for indexPtr != index {
		currentPtr = currentPtr.Next
		if currentPtr == nil {
			errMsg := "Index " + fmt.Sprintf("%d", index) + " out of range."
			return -1, errors.New(errMsg)
		}
		indexPtr++
	}

	return currentPtr.Data, nil
}

//New generates a new linked list
func New() LList {
	return LList{Next: nil}
}

//Printer prints the values in a linked list
func Printer(l LList) {

	currentPointer := l.Next

	outString := ""

	if currentPointer == nil {
		fmt.Println("Empty List")
		return
	}

	//iterate and append each value to outstring
	for currentPointer.Next != nil {
		currentPointer = currentPointer.Next
		outString += fmt.Sprintf("%d", currentPointer.Data) + " | "
	}

	outString = outString[:len(outString)-1]

	fmt.Printf(outString)
}
