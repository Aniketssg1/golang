package main

import (
	"My_Stack/my_stack"
	"fmt"
)

func main() {
	st, err := my_stack.New(3)
	if err != nil {
		fmt.Print("Error: ", err)
		return
	}
	err = st.Push(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = st.Push(2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = st.PrintStack()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
