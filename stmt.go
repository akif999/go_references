package main

import (
	"fmt"
	"go/ast"
)

func parseStatement(stmt ast.Stmt) (variableReferences, error) {
	switch stmt.(type) {
	case *ast.BadStmt:
		// Notice: Since an error occurs in the package processing that acquires Go's AST, there is no error here actually.
		return variableReferences{}, fmt.Errorf("input files has bad declaration at line %d", getLine(fset, stmt.(*ast.BadStmt).From))
	case *ast.DeclStmt:
		_, _ = parseDecl(stmt.(*ast.DeclStmt).Decl)
	case *ast.EmptyStmt:
		// EmptyStmt has no referencing variable
	case *ast.LabeledStmt:
		parseStatement(stmt.(*ast.LabeledStmt).Stmt)
	case *ast.ExprStmt:
		parseExpression(stmt.(*ast.ExprStmt).X)
	case *ast.SendStmt:
		parseExpression(stmt.(*ast.SendStmt).Chan)
		parseExpression(stmt.(*ast.SendStmt).Value)
	case *ast.IncDecStmt:
		parseExpression(stmt.(*ast.IncDecStmt).X)
	case *ast.AssignStmt:
		_, _ = parseAssignStmt(stmt.(*ast.AssignStmt))
	case *ast.GoStmt:
	case *ast.DeferStmt:
	case *ast.ReturnStmt:
		for _, e := range stmt.(*ast.ReturnStmt).Results {
			_, _ = parseExpression(e)
		}
	case *ast.BranchStmt:
	case *ast.BlockStmt:
		// BlockStmt needs to recurse (BlockStmt)
		parseBlockStmt(stmt.(*ast.BlockStmt))
	case *ast.IfStmt:
	case *ast.CaseClause:
	case *ast.SwitchStmt:
	case *ast.TypeSwitchStmt:
	case *ast.CommClause:
	case *ast.SelectStmt:
	case *ast.ForStmt:
	case *ast.RangeStmt:
		// EmptyStmt has no referencing variable
	default:
		return variableReferences{}, fmt.Errorf("invalid ast element: %+v", stmt)
	}

	return variableReferences{}, nil
}

func parseBlockStmt(block *ast.BlockStmt) (variableReferences, error) {
	var result variableReferences
	for _, stmt := range block.List {
		parseStatement(stmt)
	}
	return result, nil
}

func parseAssignStmt(stmt *ast.AssignStmt) (variableReferences, error) {
	// fmt.Printf("Lhs: %#v\n", stmt.Lhs)
	// fmt.Printf("TokPos: %#v\n", stmt.TokPos)
	// fmt.Printf("Tok: %#v\n", stmt.Tok)
	// fmt.Printf("Rhs: %#v\n", stmt.Rhs)

	hs := append(stmt.Lhs, stmt.Rhs...)

	for _, h := range hs {
		_, _ = parseExpression(h)
	}
	return variableReferences{}, nil
}
