package references

import (
	"fmt"
	"go/ast"
	"go/token"
)

func (v *VariableReferences) parseExpression(expr ast.Expr) error {
	// TODO* is necessary checking nil ??
	switch expr.(type) {
	case *ast.BadExpr:
		// Notice: Since an error occurs in the package processing that acquires Go's AST, there is no error here actually.
		return fmt.Errorf("input files has bad expression at line %d", getLine(fset, expr.(*ast.BadExpr).From))
	case *ast.Ident:
		v.parseIdent(expr.(*ast.Ident))
	case *ast.Ellipsis:
		v.parseExpression(expr.(*ast.Ellipsis).Elt)
	case *ast.BasicLit:
		// BasicLit is not variable
	case *ast.FuncLit:
		v.parseStatement(expr.(*ast.FuncLit).Body)
	case *ast.CompositeLit:
		for _, e := range expr.(*ast.CompositeLit).Elts {
			v.parseExpression(e)
		}
	case *ast.ParenExpr:
		v.parseExpression(expr.(*ast.ParenExpr).X)
	case *ast.SelectorExpr:
		v.parseExpression(expr.(*ast.SelectorExpr).X)
		v.parseIdent(expr.(*ast.SelectorExpr).Sel)
	case *ast.IndexExpr:
		v.parseExpression(expr.(*ast.IndexExpr).X)
		v.parseExpression(expr.(*ast.IndexExpr).Index)
	case *ast.IndexListExpr:
		v.parseExpression(expr.(*ast.IndexListExpr).X)
		for _, e := range expr.(*ast.IndexListExpr).Indices {
			v.parseExpression(e)
		}
	case *ast.SliceExpr:
		v.parseExpression(expr.(*ast.SliceExpr).X)
		v.parseExpression(expr.(*ast.SliceExpr).Low)
		v.parseExpression(expr.(*ast.SliceExpr).High)
		v.parseExpression(expr.(*ast.SliceExpr).Max)
	case *ast.TypeAssertExpr:
		v.parseExpression(expr.(*ast.TypeAssertExpr).X)
		v.parseExpression(expr.(*ast.TypeAssertExpr).Type)
	case *ast.CallExpr:
		v.parseExpression(expr.(*ast.CallExpr).Fun)
		for _, e := range expr.(*ast.CallExpr).Args {
			v.parseExpression(e)
		}
	case *ast.StarExpr:
		v.parseExpression(expr.(*ast.StarExpr).X)
	case *ast.UnaryExpr:
		v.parseExpression(expr.(*ast.UnaryExpr).X)
	case *ast.BinaryExpr:
		v.parseExpression(expr.(*ast.BinaryExpr).X)
		v.parseExpression(expr.(*ast.BinaryExpr).Y)
	case *ast.KeyValueExpr:
		v.parseExpression(expr.(*ast.KeyValueExpr).Key)
		v.parseExpression(expr.(*ast.KeyValueExpr).Value)
	default:
		return fmt.Errorf("invalid ast element: %+v", expr)
	}

	return nil
}

func (v *VariableReferences) parseIdent(ident *ast.Ident) {
	if !isBuiltin(ident.Name) {
		// fmt.Printf("  %s: %d\n", ident.Name, getLine(fset, ident.NamePos))
		v.Refs = append(v.Refs, VariableReference{Name: ident.Name, Row: getLine(fset, ident.NamePos)})
	}
}

func getLine(fSet *token.FileSet, pos token.Pos) int {
	ff := fSet.File(pos)
	return ff.Line(pos)
}
