package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

func parseExpression(expr ast.Expr) (variableReference, error) {
	// TODO* is necessary checking nil ??
	switch expr.(type) {
	case *ast.BadExpr:
		// Notice: Since an error occurs in the package processing that acquires Go's AST, there is no error here actually.
		return variableReference{}, fmt.Errorf("input files has bad expression at line %d", getLine(fset, expr.(*ast.BadExpr).From))
	case *ast.Ident:
		parseIdent(expr.(*ast.Ident))
	case *ast.Ellipsis:
		parseExpression(expr.(*ast.Ellipsis).Elt)
	case *ast.BasicLit:
		// BasicLit is not variable
	case *ast.FuncLit:
		_, _ = parseStatement(expr.(*ast.FuncLit).Body)
	case *ast.CompositeLit:
		for _, e := range expr.(*ast.CompositeLit).Elts {
			_, _ = parseExpression(e)
		}
	case *ast.ParenExpr:
		parseExpression(expr.(*ast.ParenExpr).X)
	case *ast.SelectorExpr:
		parseExpression(expr.(*ast.SelectorExpr).X)
		parseIdent(expr.(*ast.SelectorExpr).Sel)
	case *ast.IndexExpr:
		parseExpression(expr.(*ast.IndexExpr).X)
		parseExpression(expr.(*ast.IndexExpr).Index)
	case *ast.IndexListExpr:
		parseExpression(expr.(*ast.IndexListExpr).X)
		for _, e := range expr.(*ast.IndexListExpr).Indices {
			_, _ = parseExpression(e)
		}
	case *ast.SliceExpr:
		parseExpression(expr.(*ast.SliceExpr).X)
		parseExpression(expr.(*ast.SliceExpr).Low)
		parseExpression(expr.(*ast.SliceExpr).High)
		parseExpression(expr.(*ast.SliceExpr).Max)
	case *ast.TypeAssertExpr:
		parseExpression(expr.(*ast.TypeAssertExpr).X)
		parseExpression(expr.(*ast.TypeAssertExpr).Type)
	case *ast.CallExpr:
		parseExpression(expr.(*ast.CallExpr).Fun)
		for _, e := range expr.(*ast.CallExpr).Args {
			_, _ = parseExpression(e)
		}
	case *ast.StarExpr:
		parseExpression(expr.(*ast.StarExpr).X)
	case *ast.UnaryExpr:
		parseExpression(expr.(*ast.UnaryExpr).X)
	case *ast.BinaryExpr:
		parseExpression(expr.(*ast.BinaryExpr).X)
		parseExpression(expr.(*ast.BinaryExpr).Y)
	case *ast.KeyValueExpr:
		parseExpression(expr.(*ast.KeyValueExpr).Key)
		parseExpression(expr.(*ast.KeyValueExpr).Value)
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
