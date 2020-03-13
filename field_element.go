package main

import (
	"errors"
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)

// FieldElement struct representation of a field element
type FieldElement struct {
	num   *big.Int
	prime *big.Int
}

func (f *FieldElement) print() {
	fmt.Printf("(%.f, %.f)", f.num, f.prime)
}

// NewFieldElement returns new field element
func NewFieldElement(num *big.Int, prime *big.Int) (FieldElement, error) {

	if num.Cmp(prime) > 0 || num.Cmp(zero) == -1 {
		return FieldElement{}, errors.New("Num not in field range of prime")
	}

	fe := FieldElement{
		num:   num,
		prime: prime,
	}

	return fe, nil
}

// CheckField checks if the two elements are member of same field
func CheckField(f *big.Int, fe *big.Int) bool {

	if f.Cmp(fe) == 1 {
		return true
	}

	return false
}

// IsEqual checks if two field elemnts are equal
func (f *FieldElement) IsEqual(fe FieldElement) (bool, error) {

	var field = FieldElement{
		num:   f.num,
		prime: f.prime,
	}
	if CheckField(field.prime, fe.prime) {
		return false, errors.New("Not member of same field")
	}

	if f.num != fe.num && f.prime == fe.prime {
		return false, nil
	}

	return true, nil
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
func (f *FieldElement) Sub(fe FieldElement) (FieldElement, error) {

	if !CheckField(f.prime, fe.prime) {
		return FieldElement{}, errors.New("Not members of same field")
	}

	var res, mod big.Int

	res.Sub(f.num, fe.num)
	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld, nil
}

// Mul returns the mod multiplication of two fields
func (f *FieldElement) Mul(fe FieldElement) (FieldElement, error) {

	if !CheckField(f.prime, fe.prime) {
		return FieldElement{}, errors.New("Not member of same field")
	}

	var res, mod big.Int

	res.Mul(f.num, fe.num)
	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld, nil
}

// Pow returns the mod exponent of an element
func (f *FieldElement) Pow(exp int64) (FieldElement, error) {

	var res, mod big.Int
	var e = big.NewInt(exp)

	res.Exp(f.num, e, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld, nil
}

// Div returns the mod division of two fields
func (f *FieldElement) Div(fe FieldElement) (FieldElement, error) {

	if !CheckField(f.prime, fe.prime) {
		return FieldElement{}, errors.New("Not member of same field")
	}

	var res, mod, ex, pwr big.Int

	ex.Sub(f.prime, big.NewInt(2))
	pwr.Exp(fe.num, &ex, nil)
	res.Mul(f.num, &pwr)
	mod.Mod(&res, f.prime)

	fld := FieldElement{
		num:   &mod,
		prime: f.prime,
	}

	return fld, nil
}
