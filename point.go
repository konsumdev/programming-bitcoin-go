package main

import (
	"errors"
	"fmt"
	"math"
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

	pnt := Point{
		x: x,
		y: y,
		a: a,
		b: b,
	}
	return pnt, nil
}

// CheckIfOnCurve checks if the point is on the curve
func CheckIfOnCurve(a float64, b float64, x float64, y float64) error {

	// y2 = x3 + ax + b

	var y2 = math.Pow(y, 2)
	var x3 = math.Pow(x, 3)
	var rEq = x3 + a*x + b

	if y2 != rEq {
		return errors.New("Point is not on the curve")
	}

	return nil
}
