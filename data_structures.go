package main

import "fmt"

// ========== СТЕК (STACK) ==========
type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	val := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return val, true
}

func (s *Stack) Peek() (int, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

// ========== ОЧЕРЕДЬ (QUEUE) ==========
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) Dequeue() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	val := q.items[0]
	q.items = q.items[1:]
	return val, true
}

func (q *Queue) Peek() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	return q.items[0], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

// ========== ОДНОСВЯЗНЫЙ СПИСОК (SINGLY LINKED LIST) ==========
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Insert(val int) {
	newNode := &Node{Data: val}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (l *LinkedList) InsertAtBegin(val int) {
	newNode := &Node{Data: val}
	newNode.Next = l.Head
	l.Head = newNode
}

func (l *LinkedList) Delete(val int) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Data == val {
		l.Head = l.Head.Next
		return true
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Data == val {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList) Print() {
	current := l.Head
	fmt.Print("Список: ")
	for current != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func (l *LinkedList) Search(val int) bool {
	current := l.Head
	for current != nil {
		if current.Data == val {
			return true
		}
		current = current.Next
	}
	return false
}

// ========== ПРИМЕРЫ ИСПОЛЬЗОВАНИЯ ==========
func demonstrateStack() {
	fmt.Println("\n========== СТЕК ==========")
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println("Добавили: 10, 20, 30")

	val, ok := stack.Pop()
	fmt.Printf("Извлекли: %v (ok: %v)\n", val, ok)

	peek, _ := stack.Peek()
	fmt.Printf("Вершина стека: %v\n", peek)

	fmt.Printf("Размер стека: %d\n", stack.Size())
}

func demonstrateQueue() {
	fmt.Println("\n========== ОЧЕРЕДЬ ==========")
	queue := &Queue{}

	queue.Enqueue(5)
	queue.Enqueue(15)
	queue.Enqueue(25)
	fmt.Println("Добавили в очередь: 5, 15, 25")

	val, ok := queue.Dequeue()
	fmt.Printf("Извлекли из очереди: %v (ok: %v)\n", val, ok)

	peek, _ := queue.Peek()
	fmt.Printf("Первый в очереди: %v\n", peek)

	fmt.Printf("Размер очереди: %d\n", queue.Size())
}

func demonstrateLinkedList() {
	fmt.Println("\n========== ОДНОСВЯЗНЫЙ СПИСОК ==========")
	list := &LinkedList{}

	list.Insert(100)
	list.Insert(200)
	list.Insert(300)
	list.InsertAtBegin(50)
	list.Print()

	fmt.Printf("Ищем 200: %v\n", list.Search(200))
	fmt.Printf("Ищем 999: %v\n", list.Search(999))

	list.Delete(200)
	fmt.Print("После удаления 200: ")
	list.Print()
}

func main() {
	demonstrateStack()
	demonstrateQueue()
	demonstrateLinkedList()
}
