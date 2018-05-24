package service

import (
	"math"
)

func PGTO(valor float64, taxaJuros float64, numeroParcelas int, isInicioPeriodo bool) (valorParcela float64) {
	intInicio := 0
	//Move a divida para o vencimento
	if isInicioPeriodo {
		intInicio = 1
	} else {
		intInicio = 0
	}

	valorParcela = valor * math.Pow(1+taxaJuros, float64(-1*intInicio))
	valorParcela = (valorParcela * taxaJuros) / (1 - math.Pow(1+taxaJuros, float64(-1*numeroParcelas)))
	return
}
