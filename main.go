package main

import (
	"fmt"
)

// this is the main
func main() {
	test3()
}

// test field multiplication
func test1() {
	a, _ := NewFieldElement(3, 13)
	b, _ := NewFieldElement(12, 13)
	c, _ := NewFieldElement(10, 13)

	d, _ := a.Mul(b)
	res, _ := c.IsEqual(d)
	fmt.Println(res)
}

// test field exponentation
func test2() {
	var a, _ = NewFieldElement(3, 13)

	d, err := a.Pow(-3)
	if err != nil {
		fmt.Println(err)
	}

	e, _ := a.IsEqual(d)
	fmt.Println(e)
}

// test point add
func test3() {
	var prime int64
	prime = 223
	a, _ := NewFieldElement(5, prime)
	b, _ := NewFieldElement(7, prime)
	x1, er := NewFieldElement(-1, prime)
	if er != nil {
		fmt.Println(er)
		return
	}
	y1, er := NewFieldElement(-1, prime)
	if er != nil {
		fmt.Println(er)
		return
	}
	// x2, _ := NewFieldElement(-1, prime)
	// y2, _ := NewFieldElement(-1, prime)

	x1.print()

	p1, err := NewPoint(x1, y1, a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	// p2, _ := NewPoint(x2, y2, a, b)
	// inf, _ := NewPoint(0, 0, 5, 7)

	p1.print()
	// p2.print()
	// inf.print()

	// fmt.Println(p1.Add(&inf))
	// fmt.Println(inf.Add(&p2))
	// fmt.Println(p1.Add(&p2))
}
