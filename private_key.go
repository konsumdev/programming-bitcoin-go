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
	var kInv, sFinal big.Int
	n := hexToBigInt(N)
	nField := NewFieldElement(*n)

	nMinTwo := nField.Sub(*NewFieldElement(*big.NewInt(2)))

	k := pk.deterministicK(z)
	r := G.S256RMul(*k)
	kInv.Exp(k, nMinTwo.num, n)

	zField := NewFieldElement(*z)
	s := r.point.x.Add(*zField)
	sPoint := pk.point.S256RMul(*s.num)
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
func (pk *PrivateKey) deterministicK(z *big.Int) *big.Int {

	// Generate cryptographically strong pseudo-random between 0 - z
	rNum, _ := rand.Int(rand.Reader, z)

	return rNum
}
