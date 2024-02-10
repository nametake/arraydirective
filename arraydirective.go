package arraydirective

import (
	"slices"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gqlgo/gqlanalysis"
)

const doc = "arraydirective is ..."

// NewAnalyzer is ...
func NewAnalyzer(types, directives []string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "arraydirective",
		Doc:  doc,
		Run:  run(types, directives),
	}
}

func isTargetType(types []string, t string) bool {
	if len(types) == 0 {
		return true
	}
	return slices.Contains(types, t)
}

func run(types, directives []string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
		for _, t := range pass.Schema.Types {
			if t.BuiltIn {
				continue
			}
			if t.Kind == ast.InputObject {
				for _, field := range t.Fields {
					if field != nil && field.Type != nil && field.Type.Elem != nil {
						if isTargetType(types, field.Type.Elem.NamedType) {
							for _, directive := range directives {
								if field.Directives.ForName(directive) == nil {
									pass.Reportf(field.Position, "%s has %s directive", field.Name, directive)
								}
							}
						}
					}
				}
			}
			if t.Kind == ast.Object {
				for _, field := range t.Fields {
					for _, arg := range field.Arguments {
						if arg != nil && arg.Type != nil && arg.Type.Elem != nil {
							if isTargetType(types, arg.Type.Elem.NamedType) {
								for _, directive := range directives {
									if arg.Directives.ForName(directive) == nil {
										pass.Reportf(arg.Position, "argument %s of %s has no %s directive", arg.Name, field.Name, directive)
									}
								}
							}
						}
					}
				}
			}
		}
		return nil, nil
	}
}
