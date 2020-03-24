package main

import (
	"math/big"
)

// S256Point struct representation of s256 point
type S256Point struct {
	point *Point
}

// N we'll use the string representation for the hex value of N
const N = "0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141"

// Gx the string representation for the hex value of gx
const Gx = "0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798"

// Gy the string representation for the hex value of gy
const Gy = "0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8"

// NewS256Point init a new s256 point
func NewS256Point(x, y big.Int) S256Point {

	cmp := x.Cmp(zero)
	cmpZ := y.Cmp(zero)

	if cmp == 0 || cmpZ == 0 {
		NewP := NewPoint(*inf, *inf)
		newSP := S256Point{&NewP}
		return newSP
	}

	NewP := NewPoint(x, y)
	newSP := S256Point{&NewP}
	return newSP
}

// S256RMul redux mul for s256
func (sp *S256Point) S256RMul(coef big.Int) *S256Point {

	// try convert hex string to []bytes
	decByte := hexToBigInt(N)

	var cf big.Int
	cf.Mod(&coef, decByte)

	res := sp.point.rMul(cf)

	r256 := S256Point{res}
	return &r256
}

// verify function validate signature
/**
s_inv = pow(sig.s, N - 2, N)  # <1>
u = z * s_inv % N  # <2>
v = sig.r * s_inv % N  # <3>
total = u * G + v * self  # <4>
return total.x.num == sig.r  # <5>
*/
func (sp *S256Point) verify(z, s, r *big.Int) bool {
	// var nMinTwo, zsInv, u FieldElement
	var sInv, u, v big.Int

	zField := NewFieldElement(*z)
	rField := NewFieldElement(*r)

	n := hexToBigInt(N)
	nField := NewFieldElement(*n)

	nMinTwo := nField.Sub(*NewFieldElement(*big.NewInt(2)))

	sInv.Exp(s, nMinTwo.num, n)
	sInvField := NewFieldElement(sInv)
	zsInv := sInvField.Mul(*zField)
	u.Mod(zsInv.num, n)

	rsInv := sInvField.Mul(*rField)
	v.Mod(rsInv.num, n)

	G := gValue()

	total2 := G.S256RMul(u)
	total1 := sp.S256RMul(v)

	res := total2.point.Add(total1.point)

	return res.x.num.Cmp(r) == 0

}

func gValue() *S256Point {
	xHex := hexToBigInt(Gx)
	yHex := hexToBigInt(Gy)

	G := NewS256Point(*xHex, *yHex)
	return &G
}
