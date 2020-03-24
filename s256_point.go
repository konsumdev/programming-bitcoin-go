package main

import (
	"math/big"
)

// S256Point struct representation of s256 point
type S256Point struct {
	p Point
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
		newSP := S256Point{NewP}
		return newSP
	}

	NewP := NewPoint(x, y)
	newSP := S256Point{NewP}
	return newSP
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
func (sp *S256Point) verify(z, s, r *big.Int) bool {
	var nMinTwo, sInv, zsInv, rsInv, u, v, total, total1, total2 big.Int

	n := hexToBigInt(N)
	nMinTwo.Sub(n, big.NewInt(2))

	sInv.Exp(s, &nMinTwo, n)
	zsInv.Mul(&sInv, z)
	u.Mod(&zsInv, n)

	rsInv.Mul(&sInv, r)
	v.Mod(&rsInv, n)

	G := gValue()

	total2.Mul(&u, G.p.x.num)
	total1.Mul(&v, sp.p.x.num)
	total.Add(&total2, &total1)

	return total.Cmp(r) == 0
}

func gValue() *S256Point {
	xHex := hexToBigInt(Gx)
	yHex := hexToBigInt(Gy)

	G := NewS256Point(*xHex, *yHex)
	return &G
}
