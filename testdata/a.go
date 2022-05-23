package main

import "fmt"

var gVar1 int = 1

func func1() {
	var func1LVar1 string = "1"
	func1Lvar2 := "2"

	_ = func1LVar1
	_ = func1Lvar2
}

func main() {
	var mainLVar1 int = 1
	mainLVar2 := 2
	fmt.Println(mainLVar1)
	fmt.Println(mainLVar2)
}
