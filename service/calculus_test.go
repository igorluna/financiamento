package service

import (
	"testing"
)

func TestDeveSerCapazDeArredondarDuasCasas(t *testing.T) {
	input := 2.03621
	expected := 2.04

	calculus := Calculus{}
	actual := calculus.Round(input, 2)

	if actual != expected {
		t.Error("Arredondamento efetuado incorretamente!")
	}
}
