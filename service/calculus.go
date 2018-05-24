package service

import "math"

type Calculus struct{}

func (calculus Calculus) Round(value float64, decimals int) (roundedValue float64) {

	decimalsRound := math.Pow(float64(10), float64(decimals))

	roundedValue = math.Round(value*decimalsRound) / decimalsRound

	return
}
