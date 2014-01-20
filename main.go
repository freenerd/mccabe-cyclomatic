package main

import (
	"flag"
	"fmt"
	"github.com/freenerd/mccabe-cyclomatic/extractor"
	"log"
)

var (
	file = flag.String("f", "", "a go source file to extract calls from")
	pkg  = flag.String("p", "", "a go package")
)

func main() {
	flag.Parse()

	if *file == "" && *pkg == "" {
		log.Printf("need go source file path or go package to continue\n\n")
		flag.Usage()
		return
	}

	var complexity int64
	var err error

	if *file != "" {
		complexity, err = extractor.FileImportCalls(*file, nil)
	} else if *pkg != "" {
		complexity, err = extractor.PackageImportCalls(*pkg)
	}
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(complexity)
}
