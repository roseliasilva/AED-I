package main

import (
	"fmt"
)

type List interface {
  AddLast(value int)
  AddPos(value int, pos int)
  RemoveLast()
  Remove(pos int)
  Get(pos int) int
  Set(value int, pos int)
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

//Adicionar ao final
func (list *LinkedList) AddLast(value int){
  newNode := Node{value, nil}

  //se estiver vazia, basta add na cabeca
  if list.head == nil{
    list.head = &newNode
  }else{
    aux := list.head
    
    //navega até encontrar último nó
    for aux.next != nil{
      aux = aux.next
    }
    aux.next = &newNode
  }
  list.inserted++
}

//Adicionar em posição
func (list *LinkedList) AddPos(value int, pos int){
  //verificando se a posição é valida
  if pos>=0 && pos<=list.inserted{
    newNode := Node{value, nil}

    //se a posição for a primeira basta add na cabeca
    if pos==0{
      newNode.next = list.head
      list.head = &newNode
    }else{
      aux := list.head
      //navega até encontrar último nó
      for c:=1; aux.next!=nil && c<pos; c++{
        aux = aux.next
      }
      newNode.next = aux.next
      aux.next = &newNode 
    }
    list.inserted++
  }
}

//Remover o último elemento
func (list *LinkedList) RemoveLast(){
  list.inserted--
}

//Remove em posição específica
func (list *LinkedList) Remove(pos int){
  if pos>=0 && pos<=list.inserted{
    if pos==0{
      list.head = list.head.next
    }else{
      aux := list.head
      aux_ant := aux
      
      for i:=0; i<pos; i++{
        aux_ant = aux
        aux = aux.next
      }
      aux_ant.next = aux.next
      aux.next = nil
    }
    list.inserted--
  }
}

//Pega o valor em uma posição específica
func (list *LinkedList) Get(pos int) int {
  //verifica se a posição é válida
  if pos>=0 && pos<list.inserted{
    aux := list.head

    for c:=0; c<pos; c++{
      aux = aux.next
    }
    return aux.value
  }else{
    return -1
  }
}

//Seta o valor para uma posição específica
func (list *LinkedList) Set(value int, pos int){
  if pos>=0 && pos<list.inserted{
    aux := list.head

    for c:=0; c<pos; c++{
      aux = aux.next
    }
    aux.value = value
  }
}

//Retorna o tamanho da lista
func (list *LinkedList) Size() int {
  return list.inserted
}

func main() {
  lista := LinkedList{}

  fmt.Println("Tamanho da lista: ", lista.Size())
  //adicionando elementos
  for i:=1; i<=10; i++{
    lista.AddLast(i)
  }

  fmt.Println("Tamanho da lista: ", lista.Size())

  fmt.Println("\nElementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }

  lista.AddPos(0,7)

  fmt.Println("\nElementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }
  
  lista.RemoveLast()

  fmt.Println("\nElementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }
  
  lista.Remove(2)

  fmt.Println("\nElementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }
  
  lista.Set(9,1)

  fmt.Println("\nElementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }

  fmt.Println("\nTamanho da lista: ", lista.Size())
}