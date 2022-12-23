package main

import (
	"fmt"
	"go/ast"
)

func parseDecl(decl ast.Decl) (variableReferences, error) {
	var result variableReferences
	switch decl.(type) {
	case *ast.BadDecl:
		// Notice: Since an error occurs in the package processing that acquires Go's AST, there is no error here actually.
		return variableReferences{}, fmt.Errorf("input files has bad declaration at line %d", getLine(fset, decl.(*ast.BadDecl).From))
	case *ast.GenDecl:
	// ast.GenDecl has no variable references
	case *ast.FuncDecl:
		// ast.FuncDecl has variable references
		parseFuncDecl(decl.(*ast.FuncDecl))
	default:
		return variableReferences{}, fmt.Errorf("invalid ast element: %+v", decl)
	}

	return result, nil
}

func parseFuncDecl(decl *ast.FuncDecl) (variableReferences, error) {
	var result variableReferences
	// Doc has no variable refernces
	// Recv has no variable refernces
	// Name has no variable refernces
	// Type has no variable refernces
	// Body has variable refernces
	parseStatement(decl.Body)

	return result, nil
}
