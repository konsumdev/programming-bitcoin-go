package main

import (
	"errors"
	"math"
)

type a float64
type b float64
type x float64
type y float64

// Point is the struct representation of a point on the curve
type Point struct {
	a float64
	b float64
	x float64
	y float64
}

// CheckIfOnCurve checks if the point is on the curve
func CheckIfOnCurve(a float64, b float64, x float64, y float64) error {
	var y2 = math.Pow(y, 2)
	var x3 = math.Pow(x, 3)
	var rEq = x3 + a*x + b

	if y2 != rEq {
		return errors.New("(%f , %f) is not on the curve")
	}

	return nil
}

// IsEqual checks if two points are equal
func (p *Point) IsEqual(po Point) bool {
	return ((p.x == po.x) && (p.y == po.y)) && ((p.a == po.a) && (p.b == po.b))
}
