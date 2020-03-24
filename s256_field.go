package main

import (
	"math/big"
)

// S256Field is a struct representation of an s256 field element
type S256Field struct {
	f *FieldElement
}

// NewS256Field initialize new s256 field
func NewS256Field(num big.Int) S256Field {

	fe := NewFieldElement(num)

	fld := S256Field{fe}

	return fld
}
