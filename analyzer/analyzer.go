package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "zeros",
		Doc:  "Doc: uses for find Zero Value Structs",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.AssignStmt:
				checkForAssignZeroValue(pass, n)
			case *ast.CallExpr:
				checkForAllocationWithNew(pass, n)
			case *ast.GenDecl:
				checkForVarAssignZeroValue(pass, n)
			}

			return true
		})
	}

	return nil, nil
}

func checkForAllocationWithNew(pass *analysis.Pass, be *ast.CallExpr) {
	if ident, ok := be.Fun.(*ast.Ident); ok && ident.Name == "new" && isBuiltin(ident, pass.TypesInfo) {
		pass.Report(analysis.Diagnostic{
			Pos:     be.Pos(),
			Message: "using the new found",
		})
	}
}

func checkForAssignZeroValue(pass *analysis.Pass, be *ast.AssignStmt) {
	if be.Tok != token.DEFINE && be.Tok != token.ASSIGN {
		return
	}

	for _, r := range be.Rhs {
		checkCompositeLit(pass, r)
	}
}

func checkForVarAssignZeroValue(pass *analysis.Pass, be *ast.GenDecl) {
	if be.Tok != token.VAR {
		return
	}

	for _, spec := range be.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			continue
		}

		for _, value := range valueSpec.Values {
			checkCompositeLit(pass, value)
		}
	}
}

func checkCompositeLit(pass *analysis.Pass, expr ast.Expr) {
	cl, ok := expr.(*ast.CompositeLit)
	if !ok || len(cl.Elts) > 0 {
		return
	}

	pass.Report(analysis.Diagnostic{
		Pos:     cl.Pos(),
		Message: "zero value struct is found",
	})

	return
}

func isBuiltin(ident *ast.Ident, info *types.Info) bool {
	if ident == nil || info == nil || info.ObjectOf(ident) == nil {
		return false
	}
	obj := info.ObjectOf(ident)

	return obj != nil && obj.Pkg() == nil && obj.Parent() == types.Universe
}
