package analyzer

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var (
	skipForTest bool
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
			if skipForTest {
				return true
			}

			switch n := node.(type) {
			case *ast.AssignStmt:
				checkForAssignZeroValue(pass, n)
			case *ast.CallExpr:
				checkForAllocationWithNew(pass, n)
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
	if be.Tok != token.DEFINE {
		return
	}

	for _, r := range be.Rhs {
		if ident, ok := r.(*ast.CompositeLit); !ok || len(ident.Elts) != 0 {
			return
		}
	}

	pass.Report(analysis.Diagnostic{
		Pos:     be.Pos(),
		Message: "zero value struct is found",
	})
}

func isBuiltin(ident *ast.Ident, info *types.Info) bool {
	if ident == nil || info == nil || info.ObjectOf(ident) == nil {
		return false
	}
	obj := info.ObjectOf(ident)

	return obj != nil && obj.Pkg() == nil && obj.Parent() == types.Universe
}
