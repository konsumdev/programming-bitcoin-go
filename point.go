package main

import (
	"errors"
	"math"
	"strconv"
)

// Point is the struct representation of a point on the curve
type Point struct {
	a    float64
	b    float64
	x    float64
	y    float64
	infx bool
	infy bool
}

// NewPoint creates a newpoint struct, if x,y coordonates are empty then they will be treated as infinity
func NewPoint(x string, y string, a float64, b float64) (*Point, error) {

	var infx bool // boolean flag if x coordinate is infinity; true = infinity
	var infy bool // boolean flag if y coordinate is infinity

	if x == "" || y == "" {
		return &Point{a, b, 0, 0, true, true}, nil
	}

	// try parsing numeric string
	tempx, er := strconv.ParseFloat(x, 64)
	if er != nil {
		return &Point{}, er
	}

	tempy, er := strconv.ParseFloat(y, 64)
	if er != nil {
		return &Point{}, er
	}

	// check if the given coordinates are on the curve
	err := CheckIfOnCurve(a, b, tempx, tempy)
	if err != nil {
		return &Point{}, err
	}

	infx = false
	infy = false

	pnt := Point{
		a:    a,
		b:    b,
		x:    tempx,
		y:    tempy,
		infx: infx,
		infy: infy,
	}

	return &pnt, nil
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

// CheckSameCurve validates if the evaluated points are on the same curve
func CheckSameCurve(p *Point, po *Point) error {

	if (p.a != po.a) || (p.b != po.b) {
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
	if !p.infx {
		return po, nil
	}

	// Case 0.1: other is the point at infinity, return self
	if !po.infx {
		return p, nil
	}

	// Case 1: self.x == other.x, self.y != other.y
	// Result is point at infinity
	if (p.x == po.x) && (p.y != po.y) {
		return &Point{p.a, p.b, 0, 0, true, true}, nil
	}

	// Case 2: self.x ≠ other.x
	// Formula (x3,y3)==(x1,y1)+(x2,y2)
	// s=(y2-y1)/(x2-x1)
	// x3=s**2-x1-x2
	// y3=s*(x1-x3)-y1
	if p.x != po.x {
		var s = (po.y - p.y) / (po.x - p.x)
		var x = math.Pow(s, 2) - p.x - po.x
		var y = s*(p.x-x) - po.y

		return &Point{p.a, p.b, x, y, false, false}, nil
	}

	// Case 4: if we are tangent to the vertical line,
	// we return the point at infinity
	// note instead of figuring out what 0 is for each type
	// we just use 0 * self.x
	if (p == po) && (p.y == 0*p.x) {
		return &Point{p.a, p.b, 0, 0, false, false}, nil
	}

	// Case 3: self == other
	// Formula (x3,y3)=(x1,y1)+(x1,y1)
	// s=(3*x1**2+a)/(2*y1)
	// x3=s**2-2*x1
	// y3=s*(x1-x3)-y1
	if p == po {
		var s = (3*math.Pow(p.x, 2) + p.a) / (2 * p.y)
		var x = math.Pow(s, 2) - 2*p.x
		var y = s*(p.x-x) - p.y

		return &Point{p.a, p.b, x, y, false, false}, nil
	}

	return &Point{}, errors.New("Point addition exemption")
}
