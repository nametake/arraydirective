package arraydirective_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/nametake/arraydirective"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, arraydirective.Analyzer, "a")
}
