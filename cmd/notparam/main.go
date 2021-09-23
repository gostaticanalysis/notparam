package main

import (
	"github.com/gostaticanalysis/notparam"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(notparam.Analyzer) }
