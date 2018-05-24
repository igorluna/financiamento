package service

import (
	"testing"
)

func TestDeveSerCapazDeDozeParcelas(t *testing.T) {
	var ValorCompra float64 = 200
	numeroParcelas := 2
	taxaJuros := 0.1499

	service := FinanciamentoService{}
	actual := service.PGTO(ValorCompra, numeroParcelas, taxaJuros, false)

	calculus := Calculus{}
	actual = calculus.Round(actual, 6)
	expected := calculus.Round(123.00758267826416, 6)

	if actual != expected {
		t.Errorf("Zebrou! %f", actual)
	}
}
