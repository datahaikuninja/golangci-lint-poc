package mod2_test

import (
	"testing"

	"github.com/datahaikuninja/golangci-lint-poc/mod2"
)

func TestTypo(t *testing.T) {
	got := mod2.Typo()
	want := "misspell"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
