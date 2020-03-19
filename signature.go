package main

import "fmt"

// Signature is a struct representation of a signature
type Signature struct {
	r float64
	s float64
}

// NewSignature inits a new signature
func NewSignature(r, s float64) Signature {
	return Signature{r, s}
}

func (f *Signature) print() {
	fnum := f.r
	fprime := f.s
	fmt.Printf("Signature(%.f, %.f)\n", fprime, fnum)
}
