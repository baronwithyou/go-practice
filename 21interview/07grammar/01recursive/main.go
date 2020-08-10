package main

import (
	"fmt"
)

type People struct{
	name string
}

func (p *People) String() string {
	// 会造成递归
	//return p.String()
	// return fmt.Sprint(p)

	return ""
}

func main() {
	p := &People{}

	t := map[string]People{"test": {name: "baron"}}

	fmt.Println(t)
	fmt.Print(p.String())
}
