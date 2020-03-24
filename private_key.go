package main

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey a truct representation of a private key
type PrivateKey struct {
	secret *S256Point
	point  *S256Point
}

// NewPrivateKey inits a new private key
func NewPrivateKey(secret *Point) PrivateKey {

	privK := PrivateKey{
		secret: &S256Point{secret},
		point:  gValue(),
	}

	return privK
}

func (pk *PrivateKey) sign(z *big.Int) Signature {
	var kInv, sFinal big.Int
	n := hexToBigInt(N)
	nField := NewFieldElement(*n)

	nMinTwo := nField.Sub(*NewFieldElement(*big.NewInt(2)))

	G := gValue()
	k := pk.deterministicK(z)
	r := k.point.S256RMul(*G.point.x.num)
	kInv.Exp(k.point.point.x.num, nMinTwo.num, n)

	zField := NewFieldElement(*z)
	s, _ := zField.Add(*r.point.x)
	sPoint := pk.secret.S256RMul(*s.num)
	sPoint = sPoint.S256RMul(kInv)

	sFinal.Mod(sPoint.point.x.num, n)

	nDiv := nField.Div(*NewFieldElement(*big.NewInt(2)))

	if sFinal.Cmp(nDiv.num) == 1 {
		sRet := nField.Sub(*NewFieldElement(sFinal))
		sFinal = *sRet.num
	}

	return Signature{r.point.x.num, &sFinal}
}

// TO DO, should return a digest - hash sha256
// deterministicK unique k
func (pk *PrivateKey) deterministicK(z *big.Int) PrivateKey {

	// Generate cryptographically strong pseudo-random between 0 - z
	rNum, _ := rand.Int(rand.Reader, z)
	rPoint := NewPoint(*rNum, *rNum)
	return NewPrivateKey(&rPoint)
}
