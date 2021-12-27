package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	go_build_directive_lint "github.com/tbonnardel/go-build-directive-lint"
)

func main() {
	singlechecker.Main(go_build_directive_lint.Analyzer)
}