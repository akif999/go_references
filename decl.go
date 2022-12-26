package references

import (
	"fmt"
	"go/ast"
)

func (v *VariableReferences) parseDecl(decl ast.Decl) error {
	switch decl.(type) {
	case *ast.BadDecl:
		// Notice: Since an error occurs in the package processing that acquires Go's AST, there is no error here actually.
		return fmt.Errorf("input files has bad declaration at line %d", getLine(fset, decl.(*ast.BadDecl).From))
	case *ast.GenDecl:
	// ast.GenDecl has no variable references
	case *ast.FuncDecl:
		// ast.FuncDecl has variable references
		v.parseStatement(decl.(*ast.FuncDecl).Body)
	default:
		return fmt.Errorf("invalid ast element: %+v", decl)
	}

	return nil
}
