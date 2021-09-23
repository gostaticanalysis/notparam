package notparam

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "notparam restricts not to use type parameters in declaration of  functions and types"

var Analyzer = &analysis.Analyzer{
	Name: "notparam",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if l := n.Type.TypeParams.NumFields(); l == 1 {
				pass.Reportf(n.Pos(), "%s has a type parameter", n.Name.Name)
			} else if l > 1 {
				pass.Reportf(n.Pos(), "%s has type parameters", n.Name.Name)
			}
		case *ast.GenDecl:
			if n.Tok != token.TYPE {
				return
			}
			for _, spec := range n.Specs {
				tspec, _ := spec.(*ast.TypeSpec)
				if tspec == nil {
					continue
				}

				if l := tspec.TypeParams.NumFields(); l == 1 {
					pass.Reportf(n.Pos(), "%s has a type parameter", tspec.Name.Name)
				} else if l > 1 {
					pass.Reportf(n.Pos(), "%s has type parameters", tspec.Name.Name)
				}
			}
		}
	})

	return nil, nil
}
