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
		_, _ = parseStatement(stmt.(*ast.LabeledStmt).Stmt)
	case *ast.ExprStmt:
		_, _ = parseExpression(stmt.(*ast.ExprStmt).X)
	case *ast.SendStmt:
		_, _ = parseExpression(stmt.(*ast.SendStmt).Chan)
		_, _ = parseExpression(stmt.(*ast.SendStmt).Value)
	case *ast.IncDecStmt:
		_, _ = parseExpression(stmt.(*ast.IncDecStmt).X)
	case *ast.AssignStmt:
		for _, e := range stmt.(*ast.AssignStmt).Lhs {
			_, _ = parseExpression(e)
		}
		for _, e := range stmt.(*ast.AssignStmt).Rhs {
			_, _ = parseExpression(e)
		}
	case *ast.GoStmt:
		_, _ = parseExpression(stmt.(*ast.GoStmt).Call)
	case *ast.DeferStmt:
		_, _ = parseExpression(stmt.(*ast.DeferStmt).Call)
	case *ast.ReturnStmt:
		for _, e := range stmt.(*ast.ReturnStmt).Results {
			_, _ = parseExpression(e)
		}
	case *ast.BranchStmt:
		// TODO: Label is variable??
		parseIdent(stmt.(*ast.BranchStmt).Label)
	case *ast.BlockStmt:
		for _, s := range stmt.(*ast.BlockStmt).List {
			parseStatement(s)
		}
	case *ast.IfStmt:
		_, _ = parseStatement(stmt.(*ast.IfStmt).Init)
		_, _ = parseStatement(stmt.(*ast.IfStmt).Body)
		_, _ = parseStatement(stmt.(*ast.IfStmt).Else)
	case *ast.CaseClause:
		for _, e := range stmt.(*ast.CaseClause).List {
			_, _ = parseExpression(e)
		}
		for _, s := range stmt.(*ast.CaseClause).Body {
			_, _ = parseStatement(s)
		}
	case *ast.SwitchStmt:
		_, _ = parseStatement(stmt.(*ast.SwitchStmt).Init)
		_, _ = parseStatement(stmt.(*ast.SwitchStmt).Body)
	case *ast.TypeSwitchStmt:
		_, _ = parseStatement(stmt.(*ast.TypeSwitchStmt).Init)
		_, _ = parseStatement(stmt.(*ast.TypeSwitchStmt).Assign)
		_, _ = parseStatement(stmt.(*ast.TypeSwitchStmt).Body)
	case *ast.CommClause:
		_, _ = parseStatement(stmt.(*ast.CommClause).Comm)
		for _, s := range stmt.(*ast.CommClause).Body {
			_, _ = parseStatement(s)
		}
	case *ast.SelectStmt:
		_, _ = parseStatement(stmt.(*ast.SelectStmt).Body)
	case *ast.ForStmt:
		_, _ = parseStatement(stmt.(*ast.ForStmt).Init)
		_, _ = parseStatement(stmt.(*ast.ForStmt).Post)
		_, _ = parseStatement(stmt.(*ast.ForStmt).Body)
	case *ast.RangeStmt:
		_, _ = parseExpression(stmt.(*ast.RangeStmt).Key)
		_, _ = parseExpression(stmt.(*ast.RangeStmt).Value)
		_, _ = parseStatement(stmt.(*ast.RangeStmt).Body)
	default:
		return variableReferences{}, fmt.Errorf("invalid ast element: %+v", stmt)
	}

	return variableReferences{}, nil
}
