package main

import (
	"crypto/rand"
	"fmt"
)

// this is the main
func main() {

	textPk()
}

func textPk() {
	n := hexToBigInt(N)
	rNum, _ := rand.Int(rand.Reader, n) // generate random from 0 to n
	// rm := hexToBigInt("0xe591edebdd99ccaef1ed58e43678845851b39d5e898eaae89ce45ceac8795731")
	pk := NewPrivateKey(rNum)
	z := p2256() // 2**256
	// z := hexToBigInt("0x140e23bbb6dedc5d29a203e9a4c115208b6dc92db74168beb6d4e26abd868997")
	sig := pk.sign(z)

	truth := pk.point.verify(z, sig.r, sig.s)
	fmt.Println(truth)
}

func testVerify() {
	z := hexToBigInt("0x7c076ff316692a3d7eb3c3bb0f8b1488cf72e1afcd929e29307032997a838a3d")
	r := hexToBigInt("0xeff69ef2b1bd93a66ed5219add4fb51e11a840f404876325a1e8ffe0529a2c")
	s := hexToBigInt("0xc7207fee197d27c618aea621406f6bf5ef6fca38681d82b2f06fddbdce6feab6")
	px := hexToBigInt("0x887387e452b8eacc4acfde10d9aaf7f6d9a0f975aabb10d006e4da568744d06c")
	py := hexToBigInt("0x61de6d95231cd89026e286df3b6ae4a894a3378e393e93a0f45b666329a0ae34")
	// sig := NewSignature(NewS256Field(*r), NewS256Field(*s))

	point := NewS256Point(*px, *py)
	// point.point.print()
	// fmt.Println(s)

	ver := point.verify(z, s, r)

	fmt.Println(ver)
}

func testSignature() {
	r := hexToBigInt("0x37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6")
	s := hexToBigInt("0x8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec")

	sig := NewSignature(r, s)

	sig.print()
}
