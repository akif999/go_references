package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fmt.Println("Go")
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "./testdata/a.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}
	// ast.Print(fset, f)
	for _, decl := range f.Decls {
		fmt.Printf("%#v\n", decl)
		switch decl.(type) {
		case *ast.GenDecl:
			fmt.Printf("  %#v\n", decl)
		case *ast.FuncDecl:
			fmt.Printf("  %#v\n", decl)
		default:
		}
	}
}
