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


type ArrayList struct {
	values []int
	inserted  int //posição que vai adicionar
}

//duplicar capacidade do vetor
func (list *ArrayList) doubleArray(){
  fmt.Println("Duplicando o tamanho do array...")
  curSize := len(list.values) //tamanho do vetor
  doubledValues := make([]int, 2*curSize)
	
  for i := 0; i < curSize; i++ {
    doubledValues[i] = list.values[i]
  }
  
  list.values = doubledValues  
}

//Adicionar elemento
func (list *ArrayList) AddLast(val int){
  fmt.Println("Adicionando elemento: ", val)
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
  fmt.Println("Adiconando elemento: ", val, " na posição: ", pos)
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

//Remover o último valor
func (list *ArrayList) RemoveLast(){
  fmt.Println("Removendo o último elemento")
  list.inserted--;
}

//Remover em posição específica
func (list *ArrayList) Remove(pos int){
  fmt.Println("Removendo elemento na posição: ", pos)
  //verifica se a posição é válida
  if pos >=0 && pos <= list.inserted{
    for pos < list.inserted -1 {
      list.values[pos] = list.values[pos + 1]
      pos++
    }
    list.inserted--
  }
}

//Acessa o elemento em uma posição
func (list *ArrayList) Get(pos int) int {
  if pos >=0 && pos <= list.inserted{
    return list.values[pos]
  }
  return -1
}

//Seta o valor para uma posição específica
func (list *ArrayList) Set(value int, pos int){
  fmt.Println("Setando o valor: ", value, " na posição: ", pos)
  if pos>=0 && pos <= list.inserted{
    list.values[pos] = value
  }
}

//Tamanho do Array
func (list *ArrayList) Size() int{
  //fmt.Println("Acessando o tamanho do array...")
  return list.inserted
}

//função principal
func main() {
  lista := ArrayList{values: []int{1,2,3,4,5,6,7,8, 9, 10}, inserted: 10}

  fmt.Println("Elementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }
  
  fmt.Println("\nTamanho da lista: ", lista.Size())
  lista.AddLast(1)
  lista.AddLast(3)
  lista.AddLast(3)
  lista.AddLast(4)

  fmt.Println("Elementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }

  lista.AddPos(2,7)
  lista.RemoveLast()
  lista.Remove(5)
  lista.Set(2,11)

  fmt.Println("Elementos da lista: ")
  for i:=0; i<lista.Size(); i++{
    fmt.Print(lista.Get(i), " ")
  }

  fmt.Println("Tamanho da lista: ", lista.Size())
}
