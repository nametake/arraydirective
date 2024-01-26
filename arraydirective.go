package arraydirective

import (
	"slices"
	"strings"

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

var (
	flagTypes      string
	flagDirectives string
)

func init() {
	Analyzer.Flags.StringVar(&flagTypes, "types", "", "target types")
	Analyzer.Flags.StringVar(&flagDirectives, "directives", "", "must directives")
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	types := strings.Split(flagTypes, ",")
	directives := strings.Split(flagDirectives, ",")

	for _, t := range pass.Schema.Types {
		if t.BuiltIn {
			continue
		}
		if t.Kind == ast.InputObject {
			for _, field := range t.Fields {
				if field != nil && field.Type != nil && field.Type.Elem != nil {
					if !slices.Contains(types, field.Type.Elem.NamedType) {
						continue
					}
					for _, directive := range directives {
						if field.Directives.ForName(directive) == nil {
							pass.Reportf(field.Position, "%s has %s directive", field.Name, directive)
						}
					}
				}
			}
		}
		if t.Kind == ast.Object {
			for _, field := range t.Fields {
				for _, arg := range field.Arguments {
					if arg != nil && arg.Type != nil && arg.Type.Elem != nil {
						if !slices.Contains(types, arg.Type.Elem.NamedType) {
							continue
						}
						for _, directive := range directives {
							if arg.Directives.ForName(directive) == nil {
								pass.Reportf(arg.Position, "argument %s of %s has no %s directive", arg.Name, arg.Name, directive)
							}
						}
					}
				}
			}
		}
	}
	return nil, nil
}
