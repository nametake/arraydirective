package arraydirective_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/nametake/arraydirective"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData(t)
	if err := arraydirective.Analyzer.Flags.Set("types", "ID"); err != nil {
		t.Fatal(err)
	}
	if err := arraydirective.Analyzer.Flags.Set("directives", "list"); err != nil {
		t.Fatal(err)
	}
	analysistest.Run(t, testdata, arraydirective.Analyzer, "a")
}
