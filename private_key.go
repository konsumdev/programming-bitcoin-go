package main

import "math"

// N constant
const N = 0xfffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141

// PrivateKey a truct representation of a private key
type PrivateKey struct {
	secret S256Point
	point  S256Point
}

// NewPrivateKey inits a new private key
func NewPrivateKey(sec *Point, G *Point) PrivateKey {

	privK := PrivateKey{
		secret: S256Point{*sec},
		point:  S256Point{*G},
	}

	return privK
}

func sign(pk *PrivateKey, z *S256Point, G *S256Point) Signature {

	k := pk.deterministicK(z)
	r, _ := k.point.p.x.Mul(*G.p.x)
	kInv, _ := pk.point.p.x.Pow(k.point.p.x.num.Int64())

	s, _ := r.Mul(*pk.secret.p.x)
	s1, _ := z.p.x.Add(s)
	s2, _ := s1.Add(kInv)
	s3 := math.Mod(float64(s2.num.Int64()), N)

	n2 := N / float64(2)
	if s3 > n2 {
		s3 = N - s3
	}

	return Signature{float64(r.num.Int64()), s3}
}

// TO DO, should return a digest - hash sha256
// deterministicK unique k
func (pk *PrivateKey) deterministicK(z *S256Point) PrivateKey {
	return PrivateKey{}
}
