package wf

import (
	"strings"
	"testing"
)

func TestCleanText(t *testing.T) {
	got := CleanText("Olá, este é um teste!")
	want := "olá este é um teste"

	if strings.Compare(got, want) != 0 {
		t.Errorf("deu ruim")
	}
}
