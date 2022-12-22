package main

import (
	"fmt"
	"go/ast"
)

func (r *variableReferences) parseFile(file *ast.File) error {
	// Doc has no variable refernces
	// Package has no variable refernces
	// Name has no variable refernces
	fmt.Printf("Package: %s\n", file.Name)
	// Decls has variable refernces
	parseDecls(file.Decls)
	// Scope has no variable refernces
	// Imports has no variable refernces
	// Unresolved has no variable refernces
	// TODO: is it necessary??
	// Comments has no variable refernces

	return nil
}
