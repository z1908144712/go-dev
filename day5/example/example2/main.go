package main

import (
	"fmt"
)

type Student struct {
	name string
	age int
	score float32
}

func (s *Student) init(name string, age int, score float32) {
	s.name = name
	s.age = age
	s.score = score
}

func (s Student) get() Student{
	return s
}

func main() {
	s := new(Student)
	s.init("stu", 18, 60)
	s1 := s.get()
	fmt.Println(s1)
}