package main

// S256Field is a struct representation of an s256 field element
type S256Field struct {
	f *FieldElement
}

// NewS256Field initialize new s256 field
func NewS256Field(num int64, prime int64) (S256Field, error) {

	fe, err := NewFieldElement(num, prime)
	if err != nil {
		return S256Field{}, err
	}

	fld := S256Field{&fe}

	return fld, nil
}
