package main

// S256Point struct representation of s256 point
type S256Point struct {
	p *Point
}

// NewS256Point init a new s256 point
func NewS256Point(p *Point) (S256Point, error) {

	fld := S256Point{
		p: p,
	}

	return fld, nil
}
