package main

import (
	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/nametake/arraydirective"
)

func main() { multichecker.Main(arraydirective.Analyzer) }
