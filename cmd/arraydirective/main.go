package main

import (
	"flag"
	"strings"

	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/nametake/arraydirective"
)

func main() {
	var (
		flagTypes      string
		flagDirectives string
	)

	flag.StringVar(&flagTypes, "types", "ID", "target types")
	flag.StringVar(&flagDirectives, "directives", "list", "must directives")

	types := strings.Split(flagTypes, ",")
	directives := strings.Split(flagDirectives, ",")
	multichecker.Main(arraydirective.NewAnalyzer(types, directives))
}
