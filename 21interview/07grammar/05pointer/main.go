package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

type Teacher struct{}

func (t *Teacher) Teach() {
	fmt.Println("???")
}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	s := live()
	fmt.Printf("%#v %v %p\n", s, s == nil, s)
	// if s == nil {
	// 	fmt.Println("AAAAAAA")
	// } else {
	// 	fmt.Println("BBBBBBB")
	// }

	var t *Teacher
	fmt.Printf("%#v %v %p\n", t, t == nil, t)
	// t.Teach()
}
