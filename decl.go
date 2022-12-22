package main

import (
	"fmt"
	"go/ast"
)

func parseDecls(decls []ast.Decl) (variableReferences, error) {
	var result variableReferences
	for _, decl := range decls {
		switch decl.(type) {
		case *ast.GenDecl:
		// ast.GenDecl has no variable references
		case *ast.FuncDecl:
			// ast.FuncDecl has variable references
			parseFuncDecl(decl.(*ast.FuncDecl))
		default:
			return variableReferences{}, fmt.Errorf("invalid ast element: %+v", decl)
		}
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
	parseBlockStmt(decl.Body)

	return result, nil
}