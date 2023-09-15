package main

//implementação com linked list

import (
	"fmt"
)

type Iqueue interface{
  Enqueue(value int) //empilhar
  Dequeue() int //desempilhar
  Front() int //topo
  Empty() bool
  Size() int //tamanho
} 

type Node struct{
  val int
  next *Node
}

type LinkedListQueue struct{
  front *Node
  rear *Node
  size int
}

func(queue *LinkedListQueue) Enqueue(value int){
  newNode := Node{value, nil}

  if queue.size==0{
    queue.front = &newNode
  }else{
    queue.rear.next = &newNode
  }
  queue.rear = &newNode
  queue.size++
}

func(queue *LinkedListQueue) Dequeue() int{
  if queue.size == 0{
    return -1
  }else{
    aux := queue.front.val //lastFront
    if queue.size==1{
      queue.front = nil
      queue.rear = nil
    }else{
      queue.front = queue.front.next
    }
    queue.size--
    return aux
  }
}

func(queue *LinkedListQueue) Front() int{
  if queue.size==0{
    return -1
  }else{
    return queue.front.val
  }
}

func(queue *LinkedListQueue) Empty() bool{
  return queue.size==0
}

func(queue *LinkedListQueue) Size() int{
  return queue.size
}

func main() {
	fmt.Println("Hello, World!")
}
