package main

import (
	"errors"
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)
var inf = big.NewInt(0)

var prime = pValue()

// FieldElement struct representation of a field element
type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func (f *FieldElement) print() {
	fnum := f.num.String()
	fprime := f.prime.String()
	fmt.Printf("FieldElement_%s(%s)\n", fprime, fnum)
}

// NewFieldElement returns new field element
func NewFieldElement(num big.Int) *FieldElement {

	//if num >= prime or num < 0
	// -1 x < y
	// 0 x == y
	// 1 x > y
	cmp := num.Cmp(prime)
	cmpZ := num.Cmp(zero)

	if cmp >= 0 || cmpZ == -1 {

		resStr := fmt.Sprintf("%s not in field range of prime %s", num.String(), prime.String())
		panic(resStr)
	}

	fe := FieldElement{
		num:   &num,
		prime: prime,
	}

	return &fe
}

// CheckField checks if the two elements are member of same field
func CheckField(f *big.Int, fe *big.Int) bool {

	// -1 < | 0 == | 1 >
	if f.Cmp(fe) == 0 {
		return true
	}

	return false
}

// IsEqual checks if two field elemnts are equal
func (f *FieldElement) IsEqual(fe *FieldElement) bool {

	if f.num.Cmp(fe.num) == 0 {
		return true
	}

	return false
}

// IsNotEqual checks if two field elemnts are not equal
func (f *FieldElement) IsNotEqual(fe FieldElement) (bool, error) {
	var field = FieldElement{
		num:   f.num,
		prime: f.prime,
	}
	if CheckField(field.prime, fe.prime) {
		return false, errors.New("Not member of same field")
	}

	if f.num != fe.num && f.prime == fe.prime {
		return true, nil
	}

	return false, nil
}

// Add returns the mod sum of two fields
func (f *FieldElement) Add(fe FieldElement) (FieldElement, error) {

	if !CheckField(f.prime, fe.prime) {
		return FieldElement{}, errors.New("Not members of the same field")
	}

	var res, mod big.Int

	res.Add(f.num, fe.num)
	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld, nil
}

// Sub returns the mod subtraction of two fields
func (f *FieldElement) Sub(fe FieldElement) FieldElement {

	if !CheckField(f.prime, fe.prime) {
		panic("Not members of same field")
	}

	var res, mod big.Int

	res.Sub(f.num, fe.num)
	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld
}

// Mul returns the mod multiplication of two fields
func (f *FieldElement) Mul(fe FieldElement) FieldElement {

	if !CheckField(f.prime, fe.prime) {
		panic("Not member of same field")
	}

	var res, mod big.Int

	res.Mul(f.num, fe.num)
	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld
}

// Pow returns the mod exponent of an element
func (f *FieldElement) Pow(exp big.Int) FieldElement {

	var res, n, fprime big.Int
	// var e = big.NewInt(exp)
	fprime.Sub(f.prime, big.NewInt(1))
	n.Mod(&exp, &fprime)

	res.Exp(f.num, &n, f.prime)
	// res.Mod(&res, &fprime)

	fld := FieldElement{
		num:   &res,
		prime: f.prime,
	}

	return fld
}

// Div returns the mod division of two fields
func (f *FieldElement) Div(fe FieldElement) FieldElement {

	if !CheckField(f.prime, fe.prime) {
		panic("Not member of same field")
	}

	// num = (self.num * pow(other.num, self.prime - 2, self.prime)) % self.prime
	var prime2, res, mod big.Int

	prime2.Sub(f.prime, big.NewInt(2))

	res.Exp(fe.num, &prime2, f.prime)
	res.Mul(&res, f.num)

	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld
}

// pValue generates the value of p
// p = 2**256 - 2**32 - 977
func pValue() *big.Int {
	var two256, two32, p big.Int
	two256.Exp(big.NewInt(2), big.NewInt(256), nil)
	two32.Exp(big.NewInt(2), big.NewInt(32), nil)

	p.Sub(&two256, &two32)
	p.Sub(&p, big.NewInt(977))

	return &p
}
