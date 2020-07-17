package main

import "fmt"

// Person contains characteristic of a person
type Person struct {
	name   string
	age    int8
	gender string
}

// OptionFunc ...
type OptionFunc func(*Person)

// WithName set person's name with default value
func WithName(name string) OptionFunc {
	return func(p *Person) {
		p.name = name
	}
}

func newPerson(name, gender string, age int8) *Person {
	return &Person{
		name:   name,
		age:    age,
		gender: gender,
	}
}

func newPerson2(opts ...OptionFunc) *Person {
	p := &Person{
		name:   "Unknow",
		gender: "Unknow",
		age:    1,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func main() {
	p := newPerson("baron", "male", int8(22))
	fmt.Println(p)

	p2 := newPerson2(
		WithName("baron"),
	)
	fmt.Println(p2)
}
