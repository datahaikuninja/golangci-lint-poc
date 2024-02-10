package mod1_test

import (
	"testing"

	"github.com/datahaikuninja/golangci-lint-poc/mod1"
)

func TestTypo(t *testing.T) {
	got := mod1.Typo()
	want := "mis_spell"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
