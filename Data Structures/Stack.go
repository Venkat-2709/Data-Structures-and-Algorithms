package main

import "fmt"

type Value struct {
	data interface{}
}

type Stack struct {
	/* Containing array of type Value Struct */
	stack []*Value
}

func createStack() *Stack {
	/* Creates an Empty Stack */
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	/* Checks if an stack empty or not */
	return len(s.stack) == 0
}

func (s *Stack) push(value *Value) {
	/* Pushes the element to the stack */
	s.stack = append(s.stack, value)
	fmt.Printf("\n%v is pushed to the stack.\n", value.data)
}

func (s *Stack) pop() interface{} {
	/* Pops the element from the stack */
	if len(s.stack) == 0 {
		return nil
	}
	poppedElement := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	fmt.Printf("\n%v is popped from the stack \n", poppedElement.data)
	return poppedElement.data
}

func main() {
	/* Main Function */
	stack := createStack() // Creating a new Stack.

	fmt.Println(stack.isEmpty())
	stack.push(&Value{5})
	stack.push(&Value{9})
	stack.push(&Value{12})
	stack.push(&Value{6})
	fmt.Println(stack.isEmpty())

	fmt.Println(stack.pop(), stack.pop(), stack.pop(), stack.pop())
}
