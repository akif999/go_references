package main

import (
	"go/parser"
	"go/token"
	"log"
)

type variableReferences []variableReference

type variableReference struct {
	varName string
	row     int
	col     int
}

var fset *token.FileSet = token.NewFileSet()

func run() {
	f, err := parser.ParseFile(fset, "./testdata/a.go", nil, 0)
	if err != nil {
		log.Fatalf("parser.ParseFile failed: %s\n", err)
	}
	references := variableReferences{}
	references.parseFile(f)
}
