package references

import (
	"fmt"
	"go/ast"
)

func (v *VariableReferences) parseStatement(stmt ast.Stmt) error {
	switch stmt.(type) {
	case *ast.BadStmt:
		// Notice: Since an error occurs in the package processing that acquires Go's AST, there is no error here actually.
		return fmt.Errorf("input files has bad declaration at line %d", getLine(fset, stmt.(*ast.BadStmt).From))
	case *ast.DeclStmt:
		v.parseDecl(stmt.(*ast.DeclStmt).Decl)
	case *ast.EmptyStmt:
		// EmptyStmt has no referencing variable
	case *ast.LabeledStmt:
		v.parseStatement(stmt.(*ast.LabeledStmt).Stmt)
	case *ast.ExprStmt:
		v.parseExpression(stmt.(*ast.ExprStmt).X)
	case *ast.SendStmt:
		v.parseExpression(stmt.(*ast.SendStmt).Chan)
		v.parseExpression(stmt.(*ast.SendStmt).Value)
	case *ast.IncDecStmt:
		v.parseExpression(stmt.(*ast.IncDecStmt).X)
	case *ast.AssignStmt:
		for _, e := range stmt.(*ast.AssignStmt).Lhs {
			v.parseExpression(e)
		}
		for _, e := range stmt.(*ast.AssignStmt).Rhs {
			v.parseExpression(e)
		}
	case *ast.GoStmt:
		v.parseExpression(stmt.(*ast.GoStmt).Call)
	case *ast.DeferStmt:
		v.parseExpression(stmt.(*ast.DeferStmt).Call)
	case *ast.ReturnStmt:
		for _, e := range stmt.(*ast.ReturnStmt).Results {
			v.parseExpression(e)
		}
	case *ast.BranchStmt:
		// TODO: Label is variable??
		v.parseIdent(stmt.(*ast.BranchStmt).Label)
	case *ast.BlockStmt:
		for _, s := range stmt.(*ast.BlockStmt).List {
			v.parseStatement(s)
		}
	case *ast.IfStmt:
		v.parseStatement(stmt.(*ast.IfStmt).Init)
		v.parseStatement(stmt.(*ast.IfStmt).Body)
		v.parseStatement(stmt.(*ast.IfStmt).Else)
	case *ast.CaseClause:
		for _, e := range stmt.(*ast.CaseClause).List {
			v.parseExpression(e)
		}
		for _, s := range stmt.(*ast.CaseClause).Body {
			v.parseStatement(s)
		}
	case *ast.SwitchStmt:
		v.parseStatement(stmt.(*ast.SwitchStmt).Init)
		v.parseStatement(stmt.(*ast.SwitchStmt).Body)
	case *ast.TypeSwitchStmt:
		v.parseStatement(stmt.(*ast.TypeSwitchStmt).Init)
		v.parseStatement(stmt.(*ast.TypeSwitchStmt).Assign)
		v.parseStatement(stmt.(*ast.TypeSwitchStmt).Body)
	case *ast.CommClause:
		v.parseStatement(stmt.(*ast.CommClause).Comm)
		for _, s := range stmt.(*ast.CommClause).Body {
			v.parseStatement(s)
		}
	case *ast.SelectStmt:
		v.parseStatement(stmt.(*ast.SelectStmt).Body)
	case *ast.ForStmt:
		v.parseStatement(stmt.(*ast.ForStmt).Init)
		v.parseStatement(stmt.(*ast.ForStmt).Post)
		v.parseStatement(stmt.(*ast.ForStmt).Body)
	case *ast.RangeStmt:
		v.parseExpression(stmt.(*ast.RangeStmt).Key)
		v.parseExpression(stmt.(*ast.RangeStmt).Value)
		v.parseStatement(stmt.(*ast.RangeStmt).Body)
	default:
		return fmt.Errorf("invalid ast element: %+v", stmt)
	}

	return nil
}
