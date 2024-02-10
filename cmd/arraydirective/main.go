package main

import (
	"flag"
	"strings"

	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/nametake/arraydirective"
)

func parseDirectives(directives string) []string {
	if directives == "" {
		return nil
	}
	return strings.Split(directives, ",")
}

func parseTypes(types string) []string {
	if types == "" {
		return nil
	}
	return strings.Split(types, ",")
}

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

	types := parseTypes(flagTypes)
	directives := parseDirectives(flagDirectives)
	multichecker.Main(arraydirective.NewAnalyzer(types, directives))
}
