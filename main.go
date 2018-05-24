package main

import (
	"financiamento/model"
	"time"
)

func main() {

	//INPUT - DataCompra, DataPrimeiraParcela, ValorCompra, NumeroParcelas, TaxaDeJuros
	DataCompra := time.Date(2009, 05, 24, 0, 0, 0, 0, time.UTC)
	DataPrimeiraParcela := time.Date(2009, 06, 24, 0, 0, 0, 0, time.UTC)
	ValorCompra := model.Money{"BRL", 200}
	numeroParcelas := 2
	taxaJuros := 14.99

	//OUTPUT - PMT
}
