package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	name  string
	age   int
	score float32
	next  *Student
}

func insertTail(h *Student) {
	tail := h
	for i := 0; i < 10; i++ {
		stu := &Student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}
		tail.next = stu
		tail = stu
	}
}

func insertHead(h **Student) {
	for i := 0; i < 10; i++ {
		stu := &Student{
			name:  fmt.Sprintf("stu%d", i),
			age:   rand.Intn(100),
			score: rand.Float32() * 100,
		}
		stu.next = *h
		*h = stu
	}
}

func trans(h *Student) {
	for h != nil {
		fmt.Println(*h)
		h = h.next
	}
	fmt.Println()
}

func main() {
	head := &Student{
		name:  "head",
		age:   18,
		score: 100,
	}
	insertTail(head)
	trans(head)
	insertHead(&head)
	trans(head)
}
