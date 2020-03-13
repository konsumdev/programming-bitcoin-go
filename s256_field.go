package main

// S256Field is a struct representation of an s256 field element
type S256Field struct {
	f *FieldElement
}

// NewS256Field initialize new s256 field
func NewS256Field(fe *FieldElement) (S256Field, error) {

	fld := S256Field{fe}

	return fld, nil
}
