package service

import (
	"math"
	"time"
)

type FinanciamentoSemEntrada struct {
	DataCompra     time.Time
	DataVencimento time.Time
	ValorCompra    float64
	NumeroParcelas int
	TaxaJuros      float64
	ValorEntrada   float64
}

func (financiamento FinanciamentoSemEntrada) ValorParcela() (valor float64) {
	//Move a divida nos dias da carencia
	diasCarencia := financiamento.DataVencimento.Sub(financiamento.DataCompra).Hours() / 24

	valorV1 := financiamento.ValorCompra * math.Pow(1+financiamento.TaxaJuros, diasCarencia/30)

	return PGTO(valorV1, financiamento.TaxaJuros, financiamento.NumeroParcelas, true)
}
