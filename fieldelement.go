package main

import (
	"errors"
	"fmt"
	"math"
)

type num float64
type prime float64

// FieldElement is the struct representation of an element
type FieldElement struct {
	num   float64
	prime float64
}

// print outputs the values
func (f *FieldElement) print() string {
	return fmt.Sprintf("{%f, %f}", f.num, f.prime)
}

// NewFieldElement returns new field element
func (f *FieldElement) NewFieldElement(num float64, prime float64) error {

	if num >= prime || num < 0 {
		return errors.New("Num not in field range of prime")
	}

	f.num = num
	f.prime = prime

	return nil
}

// CheckField checks if the two elements are member of same field
func CheckField(f FieldElement, fe FieldElement) bool {
	if f.prime == fe.prime {
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
	if CheckField(field, fe) {
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
	if CheckField(field, fe) {
		return false, errors.New("Not member of same field")
	}

	if f.num != fe.num && f.prime == fe.prime {
		return true, nil
	}

	return false, nil
}

// Add returns the mod sum of two fields
func (f *FieldElement) Add(fe FieldElement) (fld FieldElement, err error) {

	var field = FieldElement{
		num:   f.num,
		prime: f.prime,
	}
	if CheckField(field, fe) {
		return fld, errors.New("Not member of same field")
	}

	var res = f.num + fe.num
	var mod = math.Mod(res, f.prime)
	fld = FieldElement{
		num:   mod,
		prime: f.prime,
	}

	return fld, nil
}

// Sub returns the mod subtraction of two fields
func (f *FieldElement) Sub(fe FieldElement) (fld FieldElement, err error) {

	var field = FieldElement{
		num:   f.num,
		prime: f.prime,
	}
	if CheckField(field, fe) {
		return fld, errors.New("Not member of same field")
	}

	var res = f.num - fe.num
	var mod = math.Mod(res, f.prime)
	fld = FieldElement{
		num:   mod,
		prime: f.prime,
	}

	return fld, nil
}

// Mul returns the mod multiplication of two fields
func (f *FieldElement) Mul(fe FieldElement) (fld FieldElement, err error) {

	var field = FieldElement{
		num:   f.num,
		prime: f.prime,
	}
	if CheckField(field, fe) {
		return fld, errors.New("Not member of same field")
	}

	var res = f.num * fe.num
	var mod = math.Mod(res, f.prime)
	fld = FieldElement{
		num:   mod,
		prime: f.prime,
	}

	return fld, nil
}

// Pow returns the mod exponent of an element
func (f *FieldElement) Pow(exp float64) (fld FieldElement, err error) {

	var res = math.Pow(f.num, exp)
	var mod = math.Mod(res, f.prime)
	fld = FieldElement{
		num:   mod,
		prime: f.prime,
	}

	return fld, nil
}

// Div returns the mod division of two fields
func (f *FieldElement) Div(fe FieldElement) (fld FieldElement, err error) {

	var field = FieldElement{
		num:   f.num,
		prime: f.prime,
	}
	if CheckField(field, fe) {
		return fld, errors.New("Not member of same field")
	}

	var ex = f.prime - 2
	var pwr = math.Pow(fe.num, ex)
	var res = f.num * pwr
	var mod = math.Mod(res, f.prime)
	fld = FieldElement{
		num:   mod,
		prime: f.prime,
	}

	return fld, nil
}
