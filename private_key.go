package main

import (
	"crypto/rand"
	"math/big"
)

// G point coordinate
var G = gValue()

// PrivateKey a truct representation of a private key
type PrivateKey struct {
	secret *big.Int
	point  *S256Point
}

// NewPrivateKey inits a new private key
func NewPrivateKey(secret *big.Int) PrivateKey {

	np := G.S256RMul(*secret)

	privK := PrivateKey{
		secret: secret,
		point:  np,
	}

	return privK
}

func (pk *PrivateKey) sign(z *big.Int) Signature {
	var kInv, rs, zrs, zrsk, s, n2, ns big.Int
	n := hexToBigInt(N)
	nField := NewFieldElement(*n)

	nMinTwo := nField.Sub(*NewFieldElement(*big.NewInt(2)))

	k := pk.deterministicK(z) //hexToBigInt("0xfab33240037374b9131c54deb7264fb3836f03c754736c73cb42376c90f5d45b") //pk.deterministicK(z)
	r := G.S256RMul(*k)

	kInv.Exp(k, nMinTwo.num, n)

	rs.Mul(r.point.x.num, pk.secret)

	zrs.Add(z, &rs)

	zrsk.Mul(&zrs, &kInv)

	s.Mod(&zrsk, n)

	n2.Div(n, big.NewInt(2))

	if s.Cmp(&n2) == 1 {
		news := ns.Sub(n, &s)
		s = *news
	}

	return Signature{r.point.x.num, &s}
}

// TO DO, should return a digest - hash sha256
// deterministicK unique k
func (pk *PrivateKey) deterministicK(z *big.Int) *big.Int {

	// Generate cryptographically strong pseudo-random between 0 - z
	rNum, _ := rand.Int(rand.Reader, z)

	return rNum
}

// p2256 is 2**256
func p2256() *big.Int {
	var two256 big.Int
	two256.Exp(big.NewInt(2), big.NewInt(256), nil)

	return &two256
}
