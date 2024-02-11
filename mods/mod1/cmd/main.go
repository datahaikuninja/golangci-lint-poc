// comments
package main

import (
	"fmt"

	"github.com/datahaikuninja/golangci-lint-poc/mod1"
)

func main() {
	s := mod1.Typo()
	fmt.Printf("This is a misspelled word: %s\n", s)
}
