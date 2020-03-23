package main

import "math/big"

// S256Field is a struct representation of an s256 field element
type S256Field struct {
	f FieldElement
}

// NewS256Field initialize new s256 field
func NewS256Field(num big.Int) (S256Field, error) {
	prime := pValue()
	fe, err := NewFieldElement(num, prime)
	if err != nil {
		return S256Field{}, err
	}

	fld := S256Field{fe}

	return fld, nil
}

// p = 2**256 - 2**32 - 977
func pValue() big.Int {
	var two256, two32, p big.Int
	two256.Exp(big.NewInt(2), big.NewInt(256), nil)
	two32.Exp(big.NewInt(2), big.NewInt(32), nil)

	p.Sub(&two256, &two32)
	p.Sub(&p, big.NewInt(977))

	return p
}
