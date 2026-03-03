package main

import "fmt"

type Node struct {
	data int   
	next *Node 
}

// LinkedList - структура списка
type LinkedList struct {
	head *Node 
}

func (l *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data, next: nil}
	
	if l.head == nil {
		l.head = newNode
		return
	}
	

	current := l.head
	for current.next != nil {
		current = current.next 
	}
	
	current.next = newNode
}

func (l *LinkedList) Display() {
	current := l.head
	
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next 
	}
	fmt.Println("nil") 
}

func main() {
	list := &LinkedList{}
	
	// Добавляем элементы в конец
	list.InsertAtEnd(1) 
	list.InsertAtEnd(2) 
	list.InsertAtEnd(3) 
	
	list.Display()
}