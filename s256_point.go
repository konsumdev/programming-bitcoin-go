package main

import "math/big"

// S256Point struct representation of s256 point
type S256Point struct {
	p *Point
}

// NewS256Point init a new s256 point
func NewS256Point(x *big.Int, y *big.Int) (S256Point, error) {

	// s_a, _ := NewS256Field(a, zero)
	// s_b, _ := NewS256Field(b, zero)

	a := big.NewInt(0)
	b := big.NewInt(0)

	p, err := NewPoint(x, y, a, b)
	if err != nil {
		return S256Point{}, err
	}

	fld := S256Point{&p}

	return fld, nil
}
