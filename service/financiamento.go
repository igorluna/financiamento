package service

import (
	"math"
)

type FinanciamentoService struct {
}

func (financiamentoService FinanciamentoService) PGTO(valor float64, numeroDeParcelas int, taxaJuros float64, isInicioPeriodo bool) (valorParcela float64) {

	intInicio := 0
	if isInicioPeriodo {
		intInicio = 1
	} else {
		intInicio = 0
	}

	valorParcela = valor * math.Pow(1+taxaJuros, float64(-1*intInicio))
	valorParcela = (valorParcela * taxaJuros) / (1 - math.Pow(1+taxaJuros, float64(-1*numeroDeParcelas)))
	return
}
