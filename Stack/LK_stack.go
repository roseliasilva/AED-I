package main

import (
	"fmt"
)

type Stack interface{
  push(value int) //empilhar
  pop() //desempilhar
  top() int //topo
  size() int //tamanho
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
func (stack *LinkedListStack) push(value int){
  stack.topo = &Node{value, stack.topo}
  stack.inserted++
}


//Desempilhar os elementos da pilha
func (stack *LinkedListStack) pop(){
  if stack.inserted>0{
    //o próximo topo não é nulo???
    if stack.topo.next!=nil{
      //fazer o topo apontar para o prox topo
      stack.topo = stack.topo.next
    }
  }
}

//Retorna o topo da pilha
func (stack *LinkedListStack) top() int {
  if stack.topo != nil{
    return stack.topo.val
  }else{
    fmt.Println("não há topo")
    return 0 //lançamento de erro
  }
}

/*
//retorna a quantidade de elementos sem modificr a pilha
func (pilha *LinkedListStack) size() int {
  
}
*/

func main() {
	//fmt.Println("Hello, World!")

  stack := LinkedListStack{}

  for i:=0; i<5; i++{
    stack.push(i)
  }

  stack.pop()
  fmt.Println(stack.top())
  
}
