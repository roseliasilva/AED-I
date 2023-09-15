package main

import (
	"fmt"
)

type List interface{
  AddLast(value int)
  AddPos(value int, pos int)
	RemoveLast()
  Remove(pos int)
  Get(pos int) int
  Set(value int, pos int)
	Size() int
}

type DoublyLinkedList struct{
  head *Node2p
  tail *Node2p
  inserted int
}

type Node2p struct{
  prev *Node2p //anterior
  value int    //valor
  next *Node2p //próximo
}

func (doublyLinkedList *DoublyLinkedList) AddLast(value int){
  newNode := Node2p{nil, value, nil}
  
}

func main() {
	fmt.Println("Hello, World!")
}
