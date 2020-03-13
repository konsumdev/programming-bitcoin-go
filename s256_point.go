package main

// S256Point struct representation of s256 point
type S256Point struct {
	x *S256Field
	y *S256Field
	a *S256Field
	b *S256Field
}

// NewS256Point init a new s256 point
func NewS256Point(x *S256Field, y *S256Field, a *S256Field, b *S256Field) (S256Point, error) {

	fld := S256Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}

	return fld, nil
}
