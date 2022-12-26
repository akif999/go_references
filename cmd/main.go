package main

import (
	"fmt"
	"log"

	references "github.com/akif999/go_references"
)

func main() {
	references := references.New()
	err := references.ParseFile("../testdata/a.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(references.File)
	for _, ref := range references.Refs {
		fmt.Printf("%s: %d\n", ref.Name, ref.Row)
	}
}
