// This example file is meant to be evaluated with mccabe-cyclomatic
//  mccabe-cyclomatic -f ./example/example.go
// It should return 8 as cyclomatic number

// This example file is also used as fixture file for go test in ../extractor/

package mccabecyclomaticexamplepackage

import (
	"fmt"
)

var t interface{}

func main() {
	a := 1

	if a == 1 {
		// pass
	} else {
		// pass
	}

	for i := 0; i < 1; i++ {
		// pass
	}

	b := []int{1}
	for _, c := range b {
		b[0] = c
		break
	}

	switch 1 {
	case 1: // pass
	default: // pass
	}

	switch t.(type) {
	case int: //pass
	case string: // pass
	default: // pass
	}

	var c chan int
	select {
	case <-c: // pass
	default: // pass
	}

	fmt.Println("Nothing happened since everything above just runs through, that's good")
}
