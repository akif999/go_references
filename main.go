package main

import (
	"fmt"
	"go/ast"
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

func (r *variableReferences) parseFile(file *ast.File) error {
	// Doc has no variable refernces
	// Package has no variable refernces
	// Name has no variable refernces
	// Decls has variable refernces
	parseDecls(file.Decls)
	// Scope has no variable refernces
	// Imports has no variable refernces
	// Unresolved has no variable refernces
	// Comments has no variable refernces

	return nil
}

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

func parseBlockStmt(block *ast.BlockStmt) (variableReferences, error) {
	var result variableReferences
	for _, stmt := range block.List {
		switch stmt.(type) {
		case *ast.BlockStmt:
			// BlockStmt needs to recurse (BlockStmt)
			parseBlockStmt(stmt.(*ast.BlockStmt))
		case *ast.ExprStmt:
		case *ast.DeclStmt:
			// DeclStmt has no referencing variable
		case *ast.AssignStmt:
			parseAssignStmt(stmt.(*ast.AssignStmt))
		case *ast.ReturnStmt:
		case *ast.IncDecStmt:
		case *ast.ForStmt:
		case *ast.RangeStmt:
		case *ast.IfStmt:
		case *ast.SwitchStmt:
		case *ast.TypeSwitchStmt:
		case *ast.LabeledStmt:
		case *ast.BranchStmt:
		case *ast.DeferStmt:
		case *ast.GoStmt:
		case *ast.SelectStmt:
		case *ast.SendStmt:
		case *ast.EmptyStmt:
		}
	}
	return result, nil
}

func parseAssignStmt(stmt *ast.AssignStmt) (variableReference, error) {
	fmt.Printf("Lhs: %#v\n", stmt.Lhs)
	fmt.Printf("TokPos: %#v\n", stmt.TokPos)
	fmt.Printf("Tok: %#v\n", stmt.Tok)
	fmt.Printf("Rhs: %#v\n", stmt.Rhs)

	hs := append(stmt.Lhs, stmt.Rhs...)

	for _, h := range hs {
		switch h.(type) {
		case *ast.Ident:
			parseIdent(h.(*ast.Ident))
		case *ast.BasicLit:
		default:
			// return variableReference{}, fmt.Errorf("Invalid Expr")
		}
	}
	return variableReference{}, nil
}

func parseIdent(ident *ast.Ident) {
	fmt.Printf("  %s: %d\n", ident.Name, getLine(fset, ident.NamePos))
}

func getLine(fSet *token.FileSet, pos token.Pos) int {
	ff := fSet.File(pos)
	return ff.Line(pos)
}

var fset *token.FileSet = token.NewFileSet()

func main() {
	f, err := parser.ParseFile(fset, "./testdata/a.go", nil, 0)
	if err != nil {
		log.Fatalf("parse.ParseFile failed: %s\n", err)
	}
	references := variableReferences{}
	references.parseFile(f)
}
