package references

import (
	"go/parser"
	"go/token"
)

var fset *token.FileSet = token.NewFileSet()

type VariableReferences struct {
	File string
	Refs []VariableReference
}

type VariableReference struct {
	Name string
	Row  int
	// col     int
}

func New() *VariableReferences {
	return &VariableReferences{}
}

func (v *VariableReferences) ParseFile(filePath string) error {
	file, err := parser.ParseFile(fset, filePath, nil, 0)
	if err != nil {
		return err
	}
	// Doc has no variable refernces
	// Package has no variable refernces
	// Name has no variable refernces
	// Decls has variable refernces
	for _, d := range file.Decls {
		v.parseDecl(d)
	}
	// Scope has no variable refernces
	// Imports has no variable refernces
	// Unresolved has no variable refernces
	// TODO: is it necessary??
	// Comments has no variable refernces

	return nil
}
