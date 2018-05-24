package service

import (
	"math"
	"time"
)

type FinanciamentoComEntrada struct {
	DataCompra     time.Time
	DataVencimento time.Time
	ValorCompra    float64
	NumeroParcelas int
	TaxaJuros      float64
	ValorEntrada   float64
}

func (financiamento FinanciamentoComEntrada) ValorParcela() (valor float64) {

	valorV1 := financiamento.ValorCompra
	if financiamento.ValorEntrada != 0 {
		valorV1 -= financiamento.ValorEntrada
	} else {
		//Caso nao tenha entrada, deve gerar o PGTO aqui mesmo
		return PGTO(financiamento.ValorCompra, financiamento.TaxaJuros, financiamento.NumeroParcelas, true)
	}
	//Move a divida nos dias da carencia
	diasCarencia := financiamento.DataVencimento.Sub(financiamento.DataCompra).Hours() / 24

	valorV1 = valorV1 * math.Pow(1+financiamento.TaxaJuros, diasCarencia/30)

	return PGTO(valorV1, financiamento.TaxaJuros, financiamento.NumeroParcelas-1, true)
}
