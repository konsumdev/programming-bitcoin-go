package main

import (
	"fmt"
)

// this is the main
func main() {

	testSignature()
}

func testVerify() {
	z := hexToBigInt("0xbc62d4b80d9e36da29c16c5d4d9f11731f36052c72401a76c23c0fb5a9b74423")
	r := hexToBigInt("0x37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6")
	s := hexToBigInt("0x8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec")
	px := hexToBigInt("0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	py := hexToBigInt("0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
	// sig := NewSignature(NewS256Field(*r), NewS256Field(*s))

	point := NewS256Point(*px, *py)

	// fmt.Println(new(big.Int).Exp(s, z, nil))

	ver := point.verify(z, s, r)

	fmt.Println(ver)
}

func testSignature() {
	r := hexToBigInt("0x37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6")
	s := hexToBigInt("0x8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec")

	sig := NewSignature(r, s)

	sig.print()
}
