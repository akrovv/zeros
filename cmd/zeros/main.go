package main

import (
	"github.com/akrovv/zeros/analyzer"
	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.New()}, nil
}
