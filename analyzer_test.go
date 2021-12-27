package go_build_directive_lint

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"os"
	"path/filepath"
	"testing"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not prepare testing environnment: %v", err)
	}

	testdata := filepath.Join(wd, "testdata")
	analysistest.Run(t, testdata, Analyzer, "p")
}