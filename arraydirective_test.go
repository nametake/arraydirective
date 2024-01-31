package arraydirective_test

import (
	"testing"

	"github.com/gqlgo/gqlanalysis/analysistest"
	"github.com/nametake/arraydirective"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	type args struct {
		types      []string
		directives []string
	}
	tests := []struct {
		dir  string
		args args
	}{
		{
			dir: "arrayparams",
			args: args{
				types:      []string{"ID"},
				directives: []string{"list"},
			},
		},
	}
	testdata := analysistest.TestData(t)
	for _, tt := range tests {
		t.Run(tt.dir, func(t *testing.T) {
			analyzer := arraydirective.NewAnalyzer(tt.args.types, tt.args.directives)
			analysistest.Run(t, testdata, analyzer, tt.dir)
		})
	}
}
