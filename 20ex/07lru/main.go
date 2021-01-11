package main

import "fmt"

func main() {
	lru := NewLRU(3)

	lru.set(1, 1)
	lru.set(2, 2)
	lru.set(3, 3)

	fmt.Println(lru.get(1))

	lru.set(4, 4)
	fmt.Println(lru.get(2))
}

type node struct {
	key  int
	val  int
	prev *node
	next *node
}

type lru struct {
	capacity int
	m        map[int]*node
	head     *node
	tail     *node
}

func NewNode(key, val int) *node {
	return &node{
		key: key,
		val: val,
	}
}

func NewLRU(capacity int) *lru {
	head := NewNode(-1, -1)
	tail := NewNode(-1, -1)

	head.next = tail
	tail.prev = head

	return &lru{
		capacity: capacity,
		m:        make(map[int]*node, capacity),
		head:     head,
		tail:     tail,
	}
}

func (l *lru) set(key, val int) {
	if l.get(key) != -1 {
		l.m[key].val = val
		return
	}

	if len(l.m) == l.capacity {
		l.removeOldest()
	}

	node := NewNode(key, val)
	l.m[key] = node
	l.moveToTail(node)
}

func (l *lru) get(key int) int {
	v, ok := l.m[key]
	if !ok {
		return -1
	}

	v.prev.next = v.next
	v.next.prev = v.prev

	l.moveToTail(v)

	return v.val
}

func (l *lru) moveToTail(node *node) {
	node.next = l.tail
	node.prev = l.tail.prev
	l.tail.prev.next = node
	l.tail.prev = node
}

func (l *lru) removeOldest() {
	delete(l.m, l.head.next.key)
	l.head.next.next.prev = l.head
	l.head.next = l.head.next.next
}
