package main

import (
	"fmt"
)

type List interface {
  AddLast(value int)
  AddPos(value int, pos int)
  Update(value int, pos int)
	RemoveLast()
  Remove(pos int)
  Get(pos int) int
	Size() int
}

type Node struct{
	value int
	next *Node
}

type LinkedList struct{
	head *Node
	inserted int
}

func main() {
	fmt.Println("teste")
}