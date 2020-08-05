package main

import "testing"

const n = 10000

func TestCommonCounter(t *testing.T) {
	c := &commonCounter{} // 非线程安全

	num, duration := test(c, n)

	t.Log(num, duration)
}

func TestMutexCounter(t *testing.T) {
	m := &mutexCounter{} // 互斥锁

	num, duration := test(m, n)

	t.Log(num, duration)
}

func TestAtomicCounter(t *testing.T) {
	a := &atomicCounter{} // 原子操作

	num, duration := test(a, n)

	t.Log(num, duration)
}
