package main

//https://replit.com/@roseliasilva/Pilha-LinkedList#main.go

import (
	"fmt"
)

type Stack interface{
  Push(value int) //empilhar
  Pop() //desempilhar
  Top() int //topo
  Empty() bool
  Size() int //tamanho
} 

type Node struct{
  val int
  next *Node
}

type LinkedListStack struct{
  topo *Node
  inserted int
}

//Empilhar os elementos na pilha
func (stack *LinkedListStack) Push(value int){
  stack.topo = &Node{value, stack.topo}
  stack.inserted++
}


//Desempilhar os elementos da pilha
func (stack *LinkedListStack) Pop(){
  if stack.inserted>0{
    //o próximo topo não é nulo???
    if stack.topo.next!=nil{
      //fazer o topo apontar para o prox topo
      stack.topo = stack.topo.next
    }else{
      stack.topo = nil
    }
    stack.inserted--
  }
}

//Retorna o topo da pilha
func (stack *LinkedListStack) Top() int {
  if stack.topo != nil{
    return stack.topo.val
  }else{
    fmt.Println("não há topo")
    return 0 //lançamento de erro
  }
}

func (stack *LinkedListStack) Empty() bool{
  return stack.inserted == 0
}

//retorna a quantidade de elementos sem modificr a pilha
func (stack *LinkedListStack) Size() int {
  return stack.inserted
}

func main() {
	//fmt.Println("Hello, World!")

  stack := LinkedListStack{}

  for i:=0; i<5; i++{
    stack.Push(i)
  }

  fmt.Println("Tamanho da pilha: ", stack.Size())
  fmt.Println("Topo da pilha: ", stack.Top())
  fmt.Println("Desempilhando..")
  stack.Pop()
  fmt.Println("Topo da pilha: ", stack.Top())
  fmt.Println("Tamanho da pilha: ", stack.Size())
}