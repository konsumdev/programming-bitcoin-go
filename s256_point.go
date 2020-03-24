package main

import (
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

// N we'll use the string representation for the hex value of N
const N = "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141"

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
func (sp *S256Point) S256RMul(coef big.Int) *S256Point {

	// try convert hex string to []bytes
	decByte := hexToBigInt(N)

	var cf big.Int
	cf.Mod(&coef, decByte)

	res := sp.p.rMul(cf)

	r256 := S256Point{*res}
	return &r256
}

// verify function validate signature
// s_inv = pow(sig.s, N - 2, N)  # <1>
// u = z * s_inv % N  # <2>
// v = sig.r * s_inv % N  # <3>
// total = u * G + v * self  # <4>
// return total.x.num == sig.r  # <5>
func verify(sp *S256Point, z *S256Point, sig *Signature, G *S256Point) bool {

	n := hexToBigInt(N)
	var nMinTwo big.Int
	var sInv, u, v S256Field
	nMinTwo.Sub(n, big.NewInt(2))

	sInv.f.num.Exp(sig.s.f.num, &nMinTwo, n)
	zsInv, _ := z.p.x.Mul(sInv.f)

	u.f.num.Mod(zsInv.num, n)
	rsInv, _ := sInv.f.Mul(sig.r.f)
	v.f.num.Mod(rsInv.num, n)

	total, _ := u.f.Mul(*G.p.x)
	total1, _ := v.f.Mul(*sp.p.x)
	total, _ = total.Add(total1)

	return (total.num.Cmp(sig.r.f.num) == 1)
}
