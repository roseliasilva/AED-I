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

func main() {
	fmt.Println("Hello, World!")
}