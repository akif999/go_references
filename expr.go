package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

func parseExpression(expr ast.Expr) (variableReference, error) {
	switch expr.(type) {
	case *ast.BadExpr:
	case *ast.Ident:
		parseIdent(expr.(*ast.Ident))
	case *ast.Ellipsis:
	case *ast.BasicLit:
		// BasicLit is not variable
	case *ast.FuncLit:
	case *ast.CompositeLit:
	case *ast.ParenExpr:
	case *ast.SelectorExpr:
	case *ast.IndexExpr:
	case *ast.IndexListExpr:
	case *ast.SliceExpr:
	case *ast.TypeAssertExpr:
	case *ast.CallExpr:
	case *ast.StarExpr:
	case *ast.UnaryExpr:
	case *ast.BinaryExpr:
	case *ast.KeyValueExpr:
	default:
		return variableReference{}, fmt.Errorf("invalid ast element: %+v", expr)
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
