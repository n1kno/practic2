# practic2
Практика 2026 Задание 2

1 Написать три структуры данных:
· Стек
· Очередь
· Односвязанный список
<details>
  <summary>Нажмите, чтобы развернуть код</summary>
```go
	
	package main
	
	import "fmt"
	
	// Stack - структура стека
	type Stack struct {
		items []int
	}
	
	// Push - добавляет элемент на вершину стека
	func (s *Stack) Push(item int) {
		// append добавляет элемент в конец среза
		s.items = append(s.items, item)
	}
	
	// Pop - удаляет и возвращает элемент с вершины стека
	func (s *Stack) Pop() int {
		if len(s.items) == 0 {
			return -1 
		}
		
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		
		return item
	}
	
	// Peek - возвращает элемент с вершины не удаляя его
	func (s *Stack) Peek() int {
		if len(s.items) == 0 {
			return -1
		}
		return s.items[len(s.items)-1]
	}
	
	func main() {
		stack := &Stack{}
		
		// Добавляем элементы
		stack.Push(1) 
		stack.Push(2) 
		stack.Push(3) 
		
		// Извлекаем элементы 
		fmt.Println(stack.Pop()) 
		fmt.Println(stack.Pop()) 
		fmt.Println(stack.Pop()) 
	}
```
