package main

import (
	"fmt"
)

// Signature is a struct representation of a signature
type Signature struct {
	r S256Field
	s S256Field
}

// NewSignature inits a new signature
func NewSignature(r, s S256Field) Signature {
	return Signature{r, s}
}

func (f *Signature) print() {
	fnum := f.r
	fprime := f.s
	fmt.Printf("Signature(%.f, %.f)\n", fprime, fnum)
}
