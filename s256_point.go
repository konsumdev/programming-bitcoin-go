package main

// S256Point struct representation of s256 point
type S256Point struct {
	p *Point
}

// NewS256Point init a new s256 point
func NewS256Point(x, y int64) (S256Point, error) {

	// s_a, _ := NewS256Field(a, zero)
	// s_b, _ := NewS256Field(b, zero)

	// a := big.NewInt(0)
	// b := big.NewInt(0)

	// p, err := NewPoint(x, y, 0, 0)
	// if err != nil {
	// return S256Point{}, err
	// }

	fld := S256Point{}

	return fld, nil
}
