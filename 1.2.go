package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}


func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1 
	}
	
	
	item := q.items[0]
	
	q.items = q.items[1:]
	
	return item
}

func (q *Queue) Front() int {
	if len(q.items) == 0 {
		return -1
	}
	return q.items[0]
}

func main() {
	queue := &Queue{}
	
	// Добавляем элементы в конец очереди
	queue.Enqueue(1) 
	queue.Enqueue(2) 
	queue.Enqueue(3) 
	
	// Извлекаем элементы 
	fmt.Println(queue.Dequeue()) 
	fmt.Println(queue.Dequeue()) 
	fmt.Println(queue.Dequeue()) 
}