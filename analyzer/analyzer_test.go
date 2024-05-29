package analyzer

import (
	"fmt"
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestValidCodeAnalysis(t *testing.T) {
	analysistest.Run(t, testdataDir(), New(), "valid")
}

func TestInvalidCodeAnalysis(t *testing.T) {
	wantErrs := []string{
		"invalid/main.go:12:2: unexpected diagnostic: zero value struct is found",
		"invalid/main.go:17:7: unexpected diagnostic: using the new found",
	}
	var gotErrs []string

	analysistest.Run(lintErrors{&gotErrs}, testdataDir(), New(), "invalid")

	slices.Sort(gotErrs)
	if !slices.Equal(wantErrs, gotErrs) {
		t.Fatalf("want: %v, got: %v", wantErrs, gotErrs)
	}
}

func testdataDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(filepath.Dir(wd), "testdata")
}

type lintErrors struct {
	Msgs *[]string
}

func (a lintErrors) Errorf(format string, args ...interface{}) {
	*a.Msgs = append(*a.Msgs, fmt.Sprintf(format, args...))
}
