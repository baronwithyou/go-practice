package main

import (
	"fmt"
	"time"
)

type Project struct{}

func (p *Project) deferError() {
	if err := recover(); err != nil {
		fmt.Println("recover: ", err)
	}
}

func (p *Project)exec(c chan interface{}) {
	for v := range c {
		fmt.Printf("%v\n", v)
	}
}

func (p *Project) run(c chan interface{}) {
	defer p.deferError()

	for {
		go p.exec(c)
		time.Sleep(time.Second * 2)
	}
}

func (p *Project) Main() {
	c := make(chan interface{}, 100)
	go p.run(c)
	go func() {
		for {
			c <- "1"
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 10000000)
}

func main() {
	p := &Project{}

	p.Main()
}
