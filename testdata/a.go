package main

import "fmt"

var gVar1 int = 1

func func1() string {
	var func1LVar1 string = "1"
	func1Lvar2 := "2"

	_ = func1LVar1
	_ = func1Lvar2

	return func1LVar1
}

func main() {
	var mainLVar1 int = 1
	mainLVar2 := mainLVar1
	fmt.Println(mainLVar1)
	fmt.Println(mainLVar2)

	return mainVar2
}
