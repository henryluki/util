package util

import (
	"errors"
	"fmt"
)

// self-type Stack

type Stack []interface{}

func (stack *Stack) Push(s interface{}) {
	*stack = append(*stack, s)
}

func (stack *Stack) Pop() (interface{}, error) {
	thisStack := *stack
	if len(thisStack) == 0 {
		return nil, errors.New("can't pop because it's empty")
	}
	x := thisStack[len(thisStack)-1]
	*stack = thisStack[:len(thisStack)-1]
	return x, nil
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't get the top because it's empty")
	}
	return stack[len(stack)-1], nil
}

// common func

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}

func p(p interface{}) {
	fmt.Println(p)
}

func StackTest() {
	// var mystack Stack
	mystack := Stack{}
	mystack.Push("hello")
	mystack.Push("world")
	mystack.Push("2015")
	mystack.Push("06")
	mystack.Push("03")
	p("Info: Run util/stack.go test function ...")
	for _, k := range mystack {
		p(k)
	}
	p(mystack)
	x, _ := mystack.Pop()
	p(x)
	s, err := mystack.Top()
	checkErr(err)
	p(s)
	p(mystack)
	p("Info: Finished ...")

}
