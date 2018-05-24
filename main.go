package main

import (
	"financiamento/service"
	"time"
)

func main() {

	//* INPUT - DataCompra, DataPrimeiraParcela, ValorCompra, NumeroParcelas, TaxaDeJuros
	financiamento := service.Financiamento{
		DataCompra:          time.Date(2009, 05, 24, 0, 0, 0, 0, time.UTC),
		DataPrimeiraParcela: time.Date(2009, 06, 24, 0, 0, 0, 0, time.UTC),
		ValorCompra:         200.00,
		NumeroParcelas:      2,
		TaxaJuros:           14.99,
	}
	//TODO Move divida para o vencimento

	pgto := financiamento.PGTO(false)

	//* OUTPUT - PMT
}
