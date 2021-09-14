package main

import (
	"bytes"
	"fmt"
)

func main() {
	x, y := NewIntSet(), NewIntSet()

	x.Add(5) // 0 | 1 << 5 = 0010 0000 = 32
	x.Add(6) // 32 | 1 << 6 = 0010 0000 | 0100 0000 = 96

	y.Add(13)
	y.Add(104)

	x.UnionWith(y)
	fmt.Println(x.String())
}

type IntSet struct {
	words []uint64
}

func NewIntSet() *IntSet {
	return &IntSet{
		words: make([]uint64, 0),
	}
}

func (i *IntSet) Has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(i.words) && i.words[word]&(1<<bit) != 0
}

func (i *IntSet) Add(x int) {
	word, bit := x/64, uint64(x%64)
	for word >= len(i.words) {
		i.words = append(i.words, 0)
	}
	i.words[word] |= 1 << bit

	// fmt.Printf("%d:%d\n", x, i.words[word])
}

func (i *IntSet) UnionWith(t *IntSet) {
	for k, v := range t.words {
		if k < len(i.words) {
			i.words[k] |= v
		} else {
			i.words = append(i.words, v)
		}
	}
}

func (i *IntSet) String() string {
	var buf bytes.Buffer
	for k, v := range i.words {
		if v == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if v&(1<<uint(i)) != 0 {
				fmt.Fprintf(&buf, "%d ", 64*k+i)
			}
		}
	}
	return buf.String()
}
