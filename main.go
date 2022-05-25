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
		log.Fatalf("parse.ParseFile failed: %s\n", err)
	}
	// ast.Print(fset, f)
	for _, decl := range f.Decls {
		// fmt.Printf("%#v\n", decl)
		switch decl.(type) {
		case *ast.BadDecl:
			log.Fatalf("Bad Declaration: %#v\n", decl.(*ast.BadDecl))
		case *ast.GenDecl:
			// GenDecl has no referencing variable
		case *ast.FuncDecl:
			// fmt.Printf("  %#v\n", decl)
			fmt.Printf("func: %s\n", decl.(*ast.FuncDecl).Name)
			for _, stmt := range decl.(*ast.FuncDecl).Body.List {
				fmt.Printf("  stmt: %#v\n", stmt)
			}
		default:
		}
	}
}
