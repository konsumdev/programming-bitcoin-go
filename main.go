package main

import (
	"fmt"
)

// this is the main
func main() {
	testpointadd()
}

func testRMul() {
	var prime int64
	prime = 223
	a, erA := NewFieldElement(0, prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(7, prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	x1, erX1 := NewFieldElement(15, prime)
	if erX1 != nil {
		fmt.Println(erX1)
		return
	}
	y1, erY1 := NewFieldElement(86, prime)
	if erY1 != nil {
		fmt.Println(erY1)
		return
	}
	p1, erP1 := NewPoint(x1, y1, a, b)
	if erP1 != nil {
		fmt.Println(erP1)
		return
	}
	fmt.Print("p1 = ")
	p1.print()

	z := p1.rMul(0)
	fmt.Print("z = ")
	z.print()
}

func testSub() {
	var prime int64
	prime = 223
	a, erA := NewFieldElement(139, prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(142, prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	c, _ := a.Sub(b)
	c.print()
}

func testPow() {
	var prime int64
	prime = 19
	a, erA := NewFieldElement(7, prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, _ := a.Pow(3)
	b.print()
}

func testDiv() {
	var prime int64
	prime = 223
	a, erA := NewFieldElement(220, prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(113, prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	c, _ := a.Div(b)
	c.print()
}

// test point add
func testpointadd() {
	var prime int64
	prime = 223
	a, erA := NewFieldElement(0, prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(7, prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	x1, erX1 := NewFieldElement(15, prime)
	if erX1 != nil {
		fmt.Println(erX1)
		return
	}
	y1, erY1 := NewFieldElement(86, prime)
	if erY1 != nil {
		fmt.Println(erY1)
		return
	}

	// x2, erX2 := NewFieldElement(17, prime)
	// if erX1 != nil {
	// 	fmt.Println(erX2)
	// }
	// y2, erY2 := NewFieldElement(56, prime)
	// if erY1 != nil {
	// 	fmt.Println(erY2)
	// }

	p1, erP1 := NewPoint(x1, y1, a, b)
	if erP1 != nil {
		fmt.Println(erP1)
		return
	}
	fmt.Print("p1 = ")
	p1.print()

	p2, erP2 := NewPoint(x1, y1, a, b)
	if erP2 != nil {
		fmt.Println(erP2)
		return
	}
	fmt.Print("p2 = ")
	p2.print()

	p3, erP3 := NewPoint(x1, y1, a, b)
	if erP3 != nil {
		fmt.Println(erP3)
		return
	}
	fmt.Print("p3 = ")
	p3.print()

	// p2, erP2 := NewPoint(x2, y2, a, b)
	// if erP2 != nil {
	// 	fmt.Println(erP2)
	// 	return
	// }
	// fmt.Print("p2 = ")
	// p2.print()

	res, err := p1.Add(&p2)
	if err != nil {
		fmt.Println(err)
	}

	// res2, _ := res.Add(&p2)
	res3, _ := res.Add(&p3)

	res3.print()
}

// test field multiplication
func test1() {
	a, _ := NewFieldElement(2, 223)
	b, _ := NewFieldElement(3, 223)
	// c, _ := NewFieldElement(10, 13)

	d, _ := a.Mul(b)
	// res := c.IsEqual(&d)
	// fmt.Println(res)
	d.print()
}

// test field exponentation
func test2() {
	var a, _ = NewFieldElement(3, 13)

	d, err := a.Pow(-3)
	if err != nil {
		fmt.Println(err)
	}

	e := a.IsEqual(&d)
	fmt.Println(e)
}

// test point exist on curve
func test3() {
	var prime int64
	prime = 223 // placeholder
	a, _ := NewFieldElement(5, prime)
	b, _ := NewFieldElement(7, prime)
	x1, er := NewFieldElement(5, prime)
	if er != nil {
		fmt.Println(er)
		return
	}
	y1, er := NewFieldElement(7, prime)
	if er != nil {
		fmt.Println(er)
		return
	}

	p1, err := NewPoint(x1, y1, a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	p1.print()
}
