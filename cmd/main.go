package main

import (
	"fmt"

	references "go_githu.com/akif999/go_references"
)

func main() {
	references := references.New()
	references.ParseFile("./testdata/a.go")
	for _, ref := range references.Refs {
		fmt.Printf("  %s: %d\n", ref.Name, ref.Row)
	}
}
