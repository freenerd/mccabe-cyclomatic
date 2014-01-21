package extractor

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func FileComplexity(file string) (int64, error) {
	v, err := newVisitor(file)
	if err != nil {
		return -1, err
	}

	v.walk()

	return v.Complexity, nil
}

type visitor struct {
	fileAst    *ast.File
	fset       *token.FileSet
	Complexity int64
}

func newVisitor(file string) (*visitor, error) {
	fset := token.NewFileSet()
	fileAst, err := parser.ParseFile(fset, file, nil, parser.ParseComments)

	if err != nil {
		return nil, err
	}

	v := visitor{
		Complexity: 1,
		fileAst:    fileAst,
	}

	return &v, nil
}

func (v *visitor) walk() *visitor {
	ast.Walk(v, v.fileAst) // calls the Visit method for each ast node
	return v
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch node := node.(type) {
	case *ast.IfStmt:
		v.Complexity++
	case *ast.ForStmt:
		v.Complexity++
	case *ast.RangeStmt:
		v.Complexity++
	case *ast.CaseClause:
		if node.List != nil {
			// is a 'case' clause, but not 'default'
			v.Complexity++
		}
	case *ast.CommClause:
		if node.Comm != nil {
			// is a 'case' clause, but not 'default'
			v.Complexity++
		}
	}

	return v
}
