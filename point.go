package main

import (
	"errors"
	"fmt"
	"math/big"
)

// Point is a struct representation of a point
type Point struct {
	x *big.Int
	y *big.Int
	a *big.Int
	b *big.Int
}

func (p *Point) print() {
	a := p.a
	b := p.b
	x := p.x
	y := p.y

	fmt.Printf("(%.f, %.f)_%.f_%.f\n\n", x, y, a, b)
}

// NewPoint init a new point
func NewPoint(x *big.Int, y *big.Int, a *big.Int, b *big.Int) (Point, error) {

	// First, check if provided coordinate is on the curve
	err := CheckIfOnCurve(x, y, a, b)
	if err != nil {
		return Point{}, err
	}

	pnt := Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}
	return pnt, nil
}

// CheckIfOnCurve checks if the point is on the curve
func CheckIfOnCurve(x *big.Int, y *big.Int, a *big.Int, b *big.Int) error {

	// y2 = x3 + ax + b

	var y2, x3, reEq big.Int
	var e2 = big.NewInt(2)
	var e3 = big.NewInt(3)

	y2.Exp(y, e2, nil)
	x3.Exp(x, e3, nil)
	reEq.Mul(a, x)
	reEq.Add(&reEq, &x3)
	reEq.Add(&reEq, b)

	if y2.Cmp(&reEq) != 0 {
		return errors.New("Point is not on the curve")
	}

	return nil
}

// CheckSameCurve evaluates if points are on the same curve(field)
func CheckSameCurve(p1 *Point, p2 *Point) error {
	var p1a, p1b, p2a, p2b big.Int
	p1a = *p1.a
	p1b = *p1.b
	p2a = *p2.a
	p2b = *p2.b

	if (p1a.Cmp(&p2a) != 0) || (p1b.Cmp(&p2b) != 0) {
		return errors.New("Points are not on the same curve")
	}

	return nil
}
