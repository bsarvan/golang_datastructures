package main

import (
	"fmt"
)


//Queue structure for implementing queue
type Queue struct {
	arr []int
	size int
	head int
	tail int
	numOfElements int
}

//New method to create a new queue
func (queue *Queue) New(size int) {
	fmt.Println("Creating new queue of size ", size)
	queue.arr = make([]int, size)
	queue.head = 0
	queue.tail = 0
	queue.numOfElements = 0
	queue.size = size
}

//Enqueue method to add elments to queue
func (queue *Queue) Enqueue(val int) {
	if (len(queue.arr) == queue.numOfElements) {
		fmt.Println("Queue had reached the limit. Resizing")
		slice2 := make([]int, queue.size)
		queue.arr = append(queue.arr, slice2...)
		//Reset head and tail
		queue.head = 0
		queue.tail = queue.numOfElements
	}

	queue.arr[queue.tail] = val
	queue.tail = (queue.tail + 1) % len(queue.arr)
	queue.numOfElements++
}

//Dequeue method to dequeue elements from queue
func (queue *Queue) Dequeue() int {
	retVal := queue.arr[queue.head]
	queue.head = (queue.head + 1) % len(queue.arr)
	queue.numOfElements = queue.numOfElements - 1
	return retVal
}

//PrintStats prints Queue stats
func (queue *Queue) PrintStats() {
	fmt.Printf("Size - %d, Len - %d, Cap - %d, Head - %d, Tail - %d, NumOfElements - %d\n", 
               queue.size, len(queue.arr), cap(queue.arr), queue.head, queue.tail, queue.numOfElements)
}

//PrintQueue method to print the contents of the queue
func (queue *Queue) PrintQueue() {
	fmt.Println("Printing the contents of the queue")
	for _,val := range queue.arr {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func main() {

	fmt.Println("Program to implement queues in golang")
	var queue *Queue

	queue = &Queue{}
	queue.New(3)

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.PrintStats()
	queue.Enqueue(40)
	queue.PrintStats()
	queue.PrintQueue()
	val := queue.Dequeue()
	fmt.Printf("Value dequeued %d\n", val)
	queue.PrintQueue()
	queue.PrintStats()
}
