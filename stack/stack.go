package main

import "fmt"

//Element struct for Stack
type Element struct {
	val int
}

//Stack structure for stack
type Stack struct {
	elements []*Element
	size     int
}

//New method to create a stack
func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

//PrintStack method to print the contents of the stack
func (stack *Stack) PrintStack() {

	fmt.Printf("Contents of the stack - ")
	elements := stack.elements
	for _, elem := range elements {
		if elem != nil {
			fmt.Printf("%d ", elem.val)
		}
	}
	fmt.Printf("\n")
}

//Push method to push an element to stack
func (stack *Stack) Push(V int) {
	fmt.Printf("Push - %d\n", V)
	var elem *Element = &Element{V}
	stack.elements = append(stack.elements, elem)
	stack.size = stack.size + 1
}

//Pop method to pop out the element from stack
func (stack *Stack) Pop() int {
	if stack.size == 0 {
		fmt.Println("Stack is Empty")
		return 0
	}
	elem := stack.elements[stack.size-1]
	fmt.Printf("Stack Length - %d, ", len(stack.elements))
	length := len(stack.elements)
	stack.elements = stack.elements[:length-1]
	stack.size = stack.size - 1
	fmt.Printf("Pop - %d\n", elem.val)
	return elem.val
}

func main() {
	var stack *Stack = &Stack{}
	stack.New()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.PrintStack()

	stack.Pop()
	stack.Pop()
	stack.PrintStack()
}
