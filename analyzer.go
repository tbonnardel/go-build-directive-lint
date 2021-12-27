// +build !foo

package go_build_directive_lint

import (
	"golang.org/x/tools/go/analysis"
	"strings"
)

const oldBuildDirective = "// +build"

var Analyzer = &analysis.Analyzer{
	Name: "gobuilddirectivelint",
	Doc: "Check that build directives follow the new syntax introduced in Go 1.17.",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, commentsGroup := range f.Comments {
			if commentsGroup == nil {
				continue
			}

			for _, comment := range commentsGroup.List {
				if comment == nil {
					continue
				}

				if strings.HasPrefix(comment.Text, oldBuildDirective) {
					pass.Reportf(comment.Pos(), "build constraints should use the 'go:build' syntax")
				}
			}
		}
	}

	return nil, nil
}