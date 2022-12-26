package main

import "fmt"

func run() {
	references := New()
	references.ParseFile("./testdata/a.go")
	for _, ref := range references.Refs {
		fmt.Printf("  %s: %d\n", ref.Name, ref.Row)
	}
}
