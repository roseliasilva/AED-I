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

type ArrayList struct {
	values []int
	inserted  int //posição que vai adicionar
}

//duplicar capacidade do vetor
func (list *ArrayList) doubleArray(){
  fmt.Println("Entrou na função doubleArray()")
  curSize := len(list.values) //tamanho do vetor
  doubledValues := make([]int, 2*curSize)
	
  for i := 0; i < curSize; i++ {
    doubledValues[i] = list.values[i]
  }
  
  list.values = doubledValues  
}

//Adicionar elemento
func (list *ArrayList) AddLast(val int){
  fmt.Println("Entrou na função AddLast()")
  //verificação para duplicar o tamanho do array se necessário
  if list.inserted >= len(list.values){
    list.doubleArray()
  }
  //adiciona os elementos
  list.values[list.inserted] = val
  list.inserted++
}

//Adicona em posição específica
func (list *ArrayList) AddPos(val int, pos int){
  fmt.Println("Entrou na função AddPos()")
  //verificando se é um valor válido
  if pos >= 0 && pos <= list.inserted{
    //verificar se array está lotado
    if list.inserted >= len(list.values){
      list.doubleArray()
    }

    for i:=list.inserted; i>pos; i--{
      list.values[i] = list.values[i-1]
    }
    list.values[pos] = val
    list.inserted++
  }
}

//Atualiza em posição específica
func (list *ArrayList) Update(val int, pos int){
  fmt.Println("Entrou na função Update()")
  //verificando se é um valor válido
  if pos >= 0 && pos <= list.inserted{
    list.values[pos] = val
  }
}

//Remover o último valor
func (list *ArrayList) RemoveLast(){
  fmt.Println("Entrou na função RemoveLast()")
  list.inserted--;
}

//Remover em posição específica
func (list *ArrayList) Remove(pos int){
  fmt.Println("Entrou na função Remove()")
  //verifica se a posição é válida
  if pos >=0 && pos < list.inserted{
    for pos < list.inserted -1 {
      list.values[pos] = list.values[pos + 1]
      pos++
    }
    list.inserted--
  }
}

//Acessa o elemento em uma posição
func (list *ArrayList) Get(pos int) int {
  if pos >=0 && pos < list.inserted{
    return list.values[pos]
  }
  return -1
}

//Tamanho do Array
func (list *ArrayList) Size() int{
  return list.inserted
}

//função principal
func main() {
  lista := ArrayList{}
  lista.values = make([]int, 10)

  lista.AddLast(1)
  lista.AddLast(2)
  lista.AddLast(3)
  lista.AddLast(4)
  lista.AddLast(5)
  lista.AddLast(6)
  lista.AddLast(7)
  lista.AddLast(8)
  lista.AddLast(9)
  lista.AddLast(10)
  fmt.Println(lista)
  lista.AddPos(7,2)
  fmt.Println(lista)
  lista.Remove(2)
  fmt.Println(lista)
}
