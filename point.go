package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
)

// A variable
var A = big.NewInt(0)

// B variable
var B = big.NewInt(7)

// Point is a struct representation of a point
type Point struct {
	x *FieldElement
	y *FieldElement
	a *FieldElement
	b *FieldElement
}

// PrintPoint outputs string representation of a point
func (p *Point) print() {

	x := p.x
	y := p.y

	xx := x.num.String()
	yy := y.num.String()
	pr := prime.String()

	if xx == "0" || yy == "0" {
		fmt.Println("Point(infinity)")
		return
	}

	fmt.Printf("Point (%s, %s)_%d_%d FieldElement(%s)\n", xx, yy, A, B, pr)
}

// NewPoint inits a new field element point
func NewPoint(x, y big.Int) Point {

	pnt := Point{
		x: NewFieldElement(x),
		y: NewFieldElement(y),
		a: NewFieldElement(*A),
		b: NewFieldElement(*B),
	}

	if (x.Cmp(zero) == 0) || (y.Cmp(zero) == 0) {
		return pnt
	}

	onCurve := CheckIfOnCurve(&x, &y)
	if !onCurve {

		xx := x.String()
		yy := y.String()

		resStr := fmt.Sprintf("Point (%s, %s)_%d_%d is not on the curve", xx, yy, A, B)
		panic(resStr)
	}

	return pnt
}

// CheckIfOnCurve checks if the point is on the curve
func CheckIfOnCurve(x *big.Int, y *big.Int) bool {

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
	reEq.Mul(A, x)
	// x3 + ax
	reEq.Add(&reEq, &x3)
	// x3 + ax + b
	reEq.Add(&reEq, B)
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

	if (!p1a.IsEqual(p2a)) || (!p1b.IsEqual(p2b)) {
		return errors.New("Points are not on the same curve")
	}

	return nil
}

// Add function performs point addition on two points
func (p *Point) Add(po *Point) *Point {

	// validate if the points to be added are on the same curve
	isSameCurve := CheckSameCurve(p, po)
	if isSameCurve != nil {
		panic(isSameCurve)
	}

	// Case 0.0: self is the point at infinity, return other
	if p.x.num.Cmp(inf) == 0 {

		return po
	}

	// Case 0.1: other is the point at infinity, return self
	if po.x.num.Cmp(inf) == 0 {

		return p
	}

	// Case 1: self.x == other.x, self.y != other.y
	// Result is point at infinity
	if (p.x.num.Cmp(po.x.num) == 0) && (p.y.num.Cmp(po.y.num) != 0) {

		infa := NewFieldElement(*inf)
		infb := NewFieldElement(*inf)
		return &Point{infa, infb, NewFieldElement(*A), NewFieldElement(*B)}
	}

	// Case 2: self.x ≠ other.x
	// Formula (x3,y3)==(x1,y1)+(x2,y2)
	// s=(y2-y1)/(x2-x1)
	// x3=s**2-x1-x2
	// y3=s*(x1-x3)-y1
	if !p.x.IsEqual(po.x) {

		ss1 := po.y.Sub(*p.y)
		ss2 := po.x.Sub(*p.x)

		sDiv := ss1.Div(ss2)

		x3 := sDiv.Pow(*big.NewInt(2))
		x3 = x3.Sub(*p.x)
		x3 = x3.Sub(*po.x)

		y := p.x.Sub(x3)
		y = sDiv.Mul(y)
		y = y.Sub(*p.y)

		return &Point{&x3, &y, p.a, p.b}

	}

	// Case 3: self == other
	// Formula (x3,y3)=(x1,y1)+(x1,y1)
	// s = (3 * x1**2 + a) / (2 * y1)
	// x3 = s**2 - 2 * x1
	// y3 = s * (x1 - x3) - y1
	if p.IsEqual(*po) {

		x12 := p.x.Pow(*big.NewInt(2))
		fe3 := NewFieldElement(*big.NewInt(3))
		sNom := x12.Mul(*fe3)
		sNom, _ = sNom.Add(*p.a)

		fe2 := NewFieldElement(*big.NewInt(2))
		sDom := p.y.Mul(*fe2)

		sDiv := sNom.Div(sDom)

		x3 := sDiv.Pow(*big.NewInt(2))
		xx := p.x.Mul(*fe2)
		x3 = x3.Sub(xx)

		y := p.x.Sub(x3)
		y = sDiv.Mul(y)
		y = y.Sub(*p.y)

		return &Point{&x3, &y, p.a, p.b}
	}

	// Case 4: if we are tangent to the vertical line,
	// we return the point at infinity
	// note instead of figuring out what 0 is for each type
	// we just use 0 * self.x
	// if self == other and self.y == 0 * self.x
	// 0 * self.x is 0
	if (p.IsEqual(*po)) && (p.y.num.Cmp(zero) == 0) {

		infa := NewFieldElement(*inf)
		infb := NewFieldElement(*inf)
		return &Point{infa, infb, NewFieldElement(*A), NewFieldElement(*B)}
	}

	// Throw exemption in case point does not fall into any of the conditions
	panic("Point addition exemption: no condition fulfilled")
}

// IsEqual checks if two points are equal
func (p *Point) IsEqual(po Point) bool {

	if !p.x.IsEqual(po.x) || !p.y.IsEqual(po.y) || !p.a.IsEqual(po.a) || !p.b.IsEqual(po.b) {
		return false
	}

	return true
}

// def __rmul__(self, coefficient):
// 	coef = coefficient
// 	current = self  # <1>
// 	result = self.__class__(None, None, self.a, self.b)  # <2>
// 	while coef:
// 		if coef & 1:  # <3>
// 			result += current
// 		current += current  # <4>
// 		coef >>= 1  # <5>
// 	return result
func (p *Point) rMul(coef big.Int) *Point {
	current := p

	newPoint := NewPoint(*inf, *inf) // init to infinity
	result := &newPoint

	// return inf if coef is 0 or below
	// -1 x < y
	// 0 x == y
	// 1 x > y
	if coef.Cmp(big.NewInt(1)) == -1 {
		return result
	}

	result = &newPoint

	// return self if coef is 1
	if coef.Cmp(big.NewInt(1)) == 0 {
		return result
	}

	var coefBit big.Int
BitShift:
	for {

		// coef & 1
		// We are looking at whether the rightmost bit is a 1. If it is, then we add the value of the current bit.
		// sets z = x & y and returns z
		coefBit.And(&coef, big.NewInt(1))
		if coefBit.Cmp(big.NewInt(1)) == 0 {
			result = result.Add(current)
		}
		current = current.Add(current)

		// z = x >> n and returns z
		coef.Rsh(&coef, 1) // coef >> 1
		// end loop when all bits have been shifted
		if coef.Cmp(zero) == 0 {
			break BitShift
		}

	}

	return result
}

// hexToBigInt will parse hex string to big int
// Note: assumes there is 0x prefix of hex string
func hexToBigInt(str string) *big.Int {

	hexStr := str[2:len(str)]

	// try convert hex string to []bytes
	decByte, err := hex.DecodeString(hexStr)
	if err != nil {
		panic(err)
	}

	// []bytes to big.Int
	z := new(big.Int)
	z.SetBytes(decByte)

	return z
}
