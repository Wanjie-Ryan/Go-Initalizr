package main

import (
	"fmt"
	"src/test"
)

func main() {
	subFunc()
	test.MyFunction()
}
func subFunc(){
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World2!")
	fmt.Println("Hello, World3!")
}