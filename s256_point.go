package main

import (
	"math"
	"math/big"
)

// S256Point struct representation of s256 point
type S256Point struct {
	p Point
}

// A variable
var A = 0

// B variable
var B = 7

// NewS256Point init a new s256 point
func NewS256Point(x, y big.Int) (S256Point, error) {

	a, _ := NewS256Field(*big.NewInt(int64(A)))
	b, _ := NewS256Field(*big.NewInt(int64(B)))

	cmp := x.Cmp(zero)
	cmpZ := y.Cmp(zero)

	if cmp == 0 || cmpZ == 0 {
		x, _ := NewS256Field(*inf)
		y, _ := NewS256Field(*inf)
		NewP, _ := NewPoint(x.f, y.f, a.f, b.f)
		newSP := S256Point{NewP}
		return newSP, nil
	}

	xx, _ := NewS256Field(x)
	yy, _ := NewS256Field(y)
	NewP, _ := NewPoint(xx.f, yy.f, a.f, b.f)
	newSP := S256Point{NewP}
	return newSP, nil
}

// S256RMul redux mul for s256
func (sp *S256Point) S256RMul(coef int) *S256Point {
	var N float64
	N = 0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141

	cf := math.Mod(float64(coef), N)

	r := sp.p.rMul(int(cf))
	r256 := S256Point{*r}
	return &r256
}

// s_inv = pow(sig.s, N - 2, N)  # <1>
// u = z * s_inv % N  # <2>
// v = sig.r * s_inv % N  # <3>
// total = u * G + v * self  # <4>
// return total.x.num == sig.r  # <5>
func verify(sp *S256Point, z *S256Point, sig *Signature, G *S256Point) bool {

	n2 := N / float64(2)
	sInv := math.Pow(sig.s, n2)
	sInvF, _ := NewS256Field(*big.NewInt(int64(sInv)))
	u, _ := z.p.x.Mul(sInvF.f)
	sigRF, _ := NewS256Field(*big.NewInt(int64(sig.r)))
	v, _ := sInvF.f.Mul(sigRF.f)

	total, _ := u.Mul(*G.p.x)
	total1, _ := v.Mul(*sp.p.x)
	total, _ = total.Add(total1)

	return total.num.Int64() == int64(sig.r)
}
