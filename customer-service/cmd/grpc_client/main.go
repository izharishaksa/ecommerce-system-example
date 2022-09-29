package main

import (
	"container/list"
	"fmt"
)

func main() {
	m := make(map[int]*list.Element, 0)
	l := list.New()
	el := l.PushFront(1)
	m[1] = el

	fmt.Println(l.Len())

	el2, _ := m[1]
	l.Remove(el2)
	delete(m, el.Value.(int))

	fmt.Println(l.Len())
}
