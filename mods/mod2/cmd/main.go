package main

import (
	"fmt"

	"github.com/datahaikuninja/golangci-lint-poc/mod2"
)

func main() {
	s := mod2.Typo()
	fmt.Printf("This is a misspelled word: %s\n", s)
}
