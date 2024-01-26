package arraydirective

import (
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gqlgo/gqlanalysis"
)

const doc = "arraydirective is ..."

// Analyzer is ...
var Analyzer = &gqlanalysis.Analyzer{
	Name: "arraydirective",
	Doc:  doc,
	Run:  run,
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	for _, q := range pass.Queries {
		for _, f := range q.Fragments {
			for _, sel := range f.SelectionSet {
				switch sel := sel.(type) {
				case *ast.Field:
					if sel.Name == "name" {
						pass.Reportf(sel.Position, "NG")
					}
				}
			}
		}
	}
	return nil, nil
}
