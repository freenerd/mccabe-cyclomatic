// This example file is meant to be evaluated with mccabe-cyclomatic
//  mccabe-cyclomatic -f ./example/example.go
// It should return 7 as cyclomatic number

package mccabecyclomaticexamplepackage

import (
	"fmt"
)

var t interface{}

func main() {
  a := 1

  if a == 1 {
    // pass
  }

  for i := 0; i < 1; i++ {
    // pass
  }

  b := []int{1}
  for _,c := range b {
    c = c
    break
  }

  switch 1 {
  case 1: //pass
  }

  switch t.(type) {
  case int: //pass
  }

  fmt.Println("Nothing happened since everything above just runs through, that's good")
}
