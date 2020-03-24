package main

import (
	"fmt"
	"math/big"
)

// Signature is a struct representation of a signature
type Signature struct {
	r *big.Int
	s *big.Int
}

// NewSignature inits a new signature
func NewSignature(r, s *big.Int) Signature {
	return Signature{r, s}
}

func (f *Signature) print() {
	fnum := f.r
	fprime := f.s

	snum := fnum.String()
	sprime := fprime.String()
	fmt.Printf("Signature(%s, %s)\n", sprime, snum)
}
