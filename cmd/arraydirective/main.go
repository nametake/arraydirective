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

	flag.StringVar(&flagDirectives, "directives", "", "required: comma-separated list of directives")
	flag.StringVar(&flagTypes, "types", "", "comma-separated list of list types")

	if flag.Parse(); flagDirectives == "" {
		flag.Usage()
		return
	}

	types := strings.Split(flagTypes, ",")
	directives := strings.Split(flagDirectives, ",")
	multichecker.Main(arraydirective.NewAnalyzer(types, directives))
}
