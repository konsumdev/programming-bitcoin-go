package main

import (
	"errors"
	"fmt"
	"math/big"
)

// Point is a struct representation of a point
type Point struct {
	x *FieldElement
	y *FieldElement
	a *FieldElement
	b *FieldElement
}

// PrintPoint outputs string representation of a point
func (p *Point) print() {
	a := p.a
	b := p.b
	x := p.x
	y := p.y

	aa := a.num.String()
	bb := b.num.String()
	xx := x.num.String()
	yy := y.num.String()
	pr := a.prime.String()

	fmt.Printf("Point (%s, %s)_%s_%s FieldElement(%s)\n", xx, yy, aa, bb, pr)
}

// NewPoint inits a new field element point
func NewPoint(x, y, a, b FieldElement) (Point, error) {

	if (x.num.Cmp(zero) == 0) || (y.num.Cmp(zero) == 0) {

		pnt := Point{
			x: &x,
			y: &y,
			a: &a,
			b: &b,
		}

		return pnt, nil
	}

	onCurve := CheckIfOnCurve(x.num, y.num, a.num, b.num, a.prime)
	if !onCurve {
		aa := a.num.String()
		bb := b.num.String()
		xx := x.num.String()
		yy := y.num.String()

		resStr := fmt.Sprintf("Point (%s, %s)_%s_%s is not on the curve", xx, yy, aa, bb)
		return Point{}, errors.New(resStr)
	}

	pnt := Point{
		x: &x,
		y: &y,
		a: &a,
		b: &b,
	}

	return pnt, nil
}

// CheckIfOnCurve checks if the point is on the curve
func CheckIfOnCurve(x *big.Int, y *big.Int, a *big.Int, b *big.Int, prime *big.Int) bool {

	// y2 = x3 + ax + b
	// 18, 77
	// 36864
	// 456533 + 7

	var y2, x3, reEq, y2Mod, reEqMod big.Int
	var e2 = big.NewInt(2)
	var e3 = big.NewInt(3)

	// y2
	y2.Exp(y, e2, nil)
	y2Mod.Mod(&y2, prime)

	//x3
	x3.Exp(x, e3, nil)
	// a*x
	reEq.Mul(a, x)
	// x3 + ax
	reEq.Add(&reEq, &x3)
	// x3 + ax + b
	reEq.Add(&reEq, b)
	reEqMod.Mod(&reEq, prime)

	res := y2Mod.Cmp(&reEqMod)
	if res != 0 {
		return false
	}

	return true
}

// CheckSameCurve evaluates if points are on the same curve(field)
func CheckSameCurve(p1 *Point, p2 *Point) error {
	// var p1a, p1b, p2a, p2b big.Int
	p1a := p1.a
	p1b := p1.b
	p2a := p2.a
	p2b := p2.b

	if !p1a.IsEqual(p2a) || !p1b.IsEqual(p2b) {
		return errors.New("Points are not on the same curve")
	}

	return nil
}

// Add function performs point addition on two points
func (p *Point) Add(po *Point) (*Point, error) {

	// validate if the points to be added are on the same curve
	isSameCurve := CheckSameCurve(p, po)
	if isSameCurve != nil {
		return &Point{}, isSameCurve
	}

	// Case 0.0: self is the point at infinity, return other
	if p.x.num.Cmp(inf) == 0 {
		fmt.Println("case 0.0")
		return po, nil
	}

	// Case 0.1: other is the point at infinity, return self
	if po.x.num.Cmp(inf) == 0 {
		fmt.Println("case 0.1")
		return p, nil
	}

	// Case 1: self.x == other.x, self.y != other.y
	// Result is point at infinity
	if (p.x.num.Cmp(po.x.num) == 0) && (p.y.num.Cmp(po.y.num) != 0) {
		fmt.Println("case 1")
		infa, _ := NewFieldElement(inf.Int64(), p.a.prime.Int64())
		infb, _ := NewFieldElement(inf.Int64(), p.b.prime.Int64())
		return &Point{&infa, &infb, p.a, p.b}, nil
	}

	// Case 2: self.x ≠ other.x
	// Formula (x3,y3)==(x1,y1)+(x2,y2)
	// s=(y2-y1)/(x2-x1)
	// x3=s**2-x1-x2
	// y3=s*(x1-x3)-y1
	if !p.x.IsEqual(po.x) {

		ss1, _ := po.y.Sub(*p.y)
		ss2, _ := po.x.Sub(*p.x)

		sDiv, _ := ss1.Div(ss2)

		x3, _ := sDiv.Pow(2)
		x3, _ = x3.Sub(*p.x)
		x3, _ = x3.Sub(*po.x)

		y, _ := p.x.Sub(x3)
		y, _ = sDiv.Mul(y)
		y, _ = y.Sub(*p.y)

		return &Point{&x3, &y, p.a, p.b}, nil

	}

	// Case 3: self == other
	// Formula (x3,y3)=(x1,y1)+(x1,y1)
	// s=(3*x1**2+a)/(2*y1)
	// x3=s**2-2*x1
	// y3=s*(x1-x3)-y1
	if p.IsEqual(*po) {
		fmt.Println("case 3")
		var s, sA, sB, s2, x3, y3, x12 big.Int
		// (3*x1**2+a)
		x12.Exp(p.x.num, big.NewInt(2), nil)
		sA.Mul(&x12, big.NewInt(3))
		sA.Add(&sA, p.a.num)

		// (2*y1)
		sB.Mul(p.y.num, big.NewInt(2))

		// s=(3*x1**2+a)/(2*y1)
		s.Div(&sA, &sB)

		//x3=s**2-2*x1
		s2.Exp(&s, big.NewInt(2), nil) // s**2
		x3.Mul(big.NewInt(2), p.y.num)
		x3.Sub(&s2, &x3)

		// y3=s*(x1-x3)-y1
		y3.Sub(p.x.num, &x3)
		y3.Mul(&s, &y3)
		y3.Sub(&y3, p.y.num)

		x, _ := NewFieldElement(x3.Int64(), p.x.prime.Int64())
		y, _ := NewFieldElement(y3.Int64(), p.y.prime.Int64())

		newP, _ := NewPoint(x, y, *p.a, *p.b)
		return &newP, nil
	}

	// Case 4: if we are tangent to the vertical line,
	// we return the point at infinity
	// note instead of figuring out what 0 is for each type
	// we just use 0 * self.x
	// if self == other and self.y == 0 * self.x
	// 0 * self.x is 0
	if (p.IsEqual(*po)) && (p.y.num.Cmp(zero) == 0) {
		fmt.Println("case 4")
		infa, _ := NewFieldElement(inf.Int64(), p.a.prime.Int64())
		infb, _ := NewFieldElement(inf.Int64(), p.b.prime.Int64())
		return &Point{&infa, &infb, p.a, p.b}, nil
	}

	// Throw exemption in case point does not fall into any of the conditions
	return &Point{}, errors.New("Point addition exemption: no condition fulfilled")
}

// IsEqual checks if two points are equal
func (p *Point) IsEqual(po Point) bool {

	if !p.x.IsEqual(po.x) || !p.y.IsEqual(po.y) || !p.a.IsEqual(po.a) || !p.b.IsEqual(po.b) {
		return false
	}

	return true
}
