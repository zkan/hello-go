package sloths

import (
	"math/rand"
)

type randNumberGenerator interface {
	randomInt(max int) int
}

// our standard-library implementation is an empty
// struct whose randomInt method calls math/rand.Intn
type standardRand struct{}

func (s standardRand) randomInt(max int) int {
	return rand.Intn(max)
}

func divByRand(numerator int, r randNumberGenerator) int {
	//return numerator / int(rand.Intn(10))
	//return numerator / r.randomInt(10)
	denominator := 1 + int(r.randomInt(10))
	return numerator / denominator
}
