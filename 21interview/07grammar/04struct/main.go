package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func paseStudent() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Println(stu)

		m[stu.Name] = &stu
	}

	// 	for name, stu := range m {
	// 		fmt.Println(name, *stu)
	// 	}
}

func main() {
	paseStudent()
}
