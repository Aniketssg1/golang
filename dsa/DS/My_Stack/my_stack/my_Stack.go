package my_stack

import (
	"errors"
	"fmt"
)

type (
	Stack struct {
		top      *node
		length   int
		capacity int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

func New(capacity int) (*Stack, error) {
	if capacity <= 0 {
		return nil, errors.New("Stack Length Cannot Be Negative !!!")
	}
	return &Stack{nil, 0, capacity}, nil
}

func (this *Stack) Peek() (interface{}, error) {
	if this.length == 0 {
		return nil, errors.New("Stack Is Empty !!!")
	}
	return this.top.value, nil
}

func (this *Stack) Pop() (interface{}, error) {
	if this.length == 0 {
		return nil, errors.New("Stack Is Empty!!!")
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value, nil
}

func (this *Stack) Push(value interface{}) error {
	if this.length > this.capacity {
		return errors.New("Stack Is Full !!!")
	}
	n := &node{value, this.top}
	this.top = n
	this.length++
	return nil
}

func (this *Stack) PrintStack() error {
	if this.length == 0 {
		return errors.New("Stack Is Empty !!!")
	}
	n := this.top
	for n != nil {
		fmt.Print(n.value, " ")
		n = n.prev
	}
	fmt.Println()
	return nil
}
