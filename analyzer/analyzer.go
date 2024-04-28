package analyzer

import (
	"flag"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
)

//nolint:gochecknoglobals
var flagSet flag.FlagSet

//nolint:gochecknoglobals
var (
	skipForTest bool
)

//nolint:gochecknoinits
func init() {
	flag.BoolVar(&skipForTest, "skip_for_test", false, "skip linter for negative test")
}

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "zeros",
		Doc:   "Doc: uses for find Zero Value Structs",
		Run:   run,
		Flags: flagSet,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(n ast.Node) bool {
			be, ok := n.(*ast.AssignStmt)
			if !ok {
				return true
			}
			if be.Tok != token.DEFINE {
				return true
			}

			for _, r := range be.Rhs {
				if ident, ok := r.(*ast.CompositeLit); ok {
					if len(ident.Elts) != 0 {
						return true
					}
				} else {
					return true
				}
			}

			if skipForTest {
				return true
			}

			pass.Report(analysis.Diagnostic{
				Pos:     be.Pos(),
				Message: "Zero Value Structs found",
			})

			return true
		})
	}

	return nil, nil
}
