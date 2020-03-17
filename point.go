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

// IsEqual checks if two points are equal
func (p *Point) IsEqual(po Point) bool {
	return ((p.x == po.x) && (p.y == po.y)) && ((p.a == po.a) && (p.b == po.b))
}

// Add function performs point addition on two points
func (p *Point) Add(po *Point) (*Point, error) {

	// validate if the points to be added are on the same curve
	err := CheckSameCurve(p, po)
	if err != nil {
		return &Point{}, err
	}

	// Case 0.0: self is the point at infinity, return other
	if p.x.Cmp(zero) == 0 {
		return po, nil
	}

	// Case 0.1: other is the point at infinity, return self
	if po.x.Cmp(zero) == 0 {
		return p, nil
	}

	// Case 1: self.x == other.x, self.y != other.y
	// Result is point at infinity
	if (p.x.Cmp(po.x) == 0) && (p.y.Cmp(po.y) != 0) {
		return &Point{zero, zero, p.a, p.b}, nil
	}

	// Case 2: self.x ≠ other.x
	// Formula (x3,y3)==(x1,y1)+(x2,y2)
	// s=(y2-y1)/(x2-x1)
	// x3=s**2-x1-x2
	// y3=s*(x1-x3)-y1
	if p.x != po.x {
		var s1, s2, s, x, y big.Int

		s1.Sub(po.y, p.y)
		s2.Sub(po.x, p.x)
		s.Div(&s1, &s2)

		x.Exp(&s, big.NewInt(2), nil)
		x.Sub(&x, p.x)
		x.Sub(&x, po.x)

		y.Sub(p.x, &x)
		y.Mul(&s, &y)
		y.Sub(&y, po.y)

		return &Point{&x, &y, p.a, p.b}, nil
	}

	// Case 4: if we are tangent to the vertical line,
	// we return the point at infinity
	// note instead of figuring out what 0 is for each type
	// we just use 0 * self.x
	var rs big.Int
	if (p.IsEqual(*po)) && (p.y == rs.Mul(p.x, zero)) {
		return &Point{zero, zero, p.a, p.b}, nil
	}

	// Case 3: self == other
	// Formula (x3,y3)=(x1,y1)+(x1,y1)
	// s=(3*x1**2+a)/(2*y1)
	// x3=s**2-2*x1
	// y3=s*(x1-x3)-y1
	if p == po {
		var s, s1, s2, x3, y3 big.Int
		s.Exp(&s, big.NewInt(2), nil)
		s.Mul(&s, big.NewInt(3))
		s.Add(&s, p.a)
		s1.Mul(p.y, big.NewInt(2))
		s.Div(&s, &s1)

		x3.Mul(big.NewInt(2), p.x)
		s2.Exp(&s, big.NewInt(2), nil)
		x3.Sub(&s2, &x3)

		y3.Sub(p.x, &x3)
		y3.Mul(&s, &y3)
		y3.Sub(&y3, p.y)

		return &Point{&x3, &y3, p.a, p.b}, nil
	}

	// Throw exemption in case point does not fall into any of the conditions
	return &Point{}, errors.New("Point addition exemption")
}

// RMul scalar multiplication of a point
func (p *Point) RMul(coef int) (*Point, error) {
	var prod Point
	prod = *p
	for i := 1; i <= coef; i++ {
		prod.Add(p)
	}

	return &prod, nil
}
