package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

// this is the main
func main() {
	testHex()
}

func testHex() {
	N := "fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141"

	// const s = "48656c6c6f20476f7068657221"
	decByte, err := hex.DecodeString(N)
	if err != nil {
		log.Fatal(err)
	}

	z := new(big.Int)
	z.SetBytes(decByte)

	fmt.Printf("%s\n", z)
}

func testRMul() {

	prime := big.NewInt(223)
	a, erA := NewFieldElement(*big.NewInt(0), *prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(*big.NewInt(7), *prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	x1, erX1 := NewFieldElement(*big.NewInt(15), *prime)
	if erX1 != nil {
		fmt.Println(erX1)
		return
	}
	y1, erY1 := NewFieldElement(*big.NewInt(86), *prime)
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

	coef := big.NewInt(2)
	z := p1.rMul(*coef)
	fmt.Print("z = ")
	z.print()
}

func testSub() {
	prime := big.NewInt(223)
	a, erA := NewFieldElement(*big.NewInt(139), *prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(*big.NewInt(142), *prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	fmt.Printf("%d - %d = ", 139, 142)
	c, _ := a.Sub(b)
	c.print()
}

func testPow() {
	prime := big.NewInt(19)
	a, erA := NewFieldElement(*big.NewInt(7), *prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, _ := a.Pow(3)
	b.print()
}

func testDiv() {
	prime := big.NewInt(19)
	a, erA := NewFieldElement(*big.NewInt(220), *prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(*big.NewInt(113), *prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	c, _ := a.Div(b)
	c.print()
}

// test point add
func testpointadd() {
	prime := big.NewInt(223)
	a, erA := NewFieldElement(*big.NewInt(0), *prime)
	if erA != nil {
		fmt.Println(erA)
		return
	}
	b, erB := NewFieldElement(*big.NewInt(7), *prime)
	if erB != nil {
		fmt.Println(erB)
		return
	}
	x1, erX1 := NewFieldElement(*big.NewInt(15), *prime)
	if erX1 != nil {
		fmt.Println(erX1)
		return
	}
	y1, erY1 := NewFieldElement(*big.NewInt(86), *prime)
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
	prime := big.NewInt(223)
	a, _ := NewFieldElement(*big.NewInt(2), *prime)
	b, _ := NewFieldElement(*big.NewInt(3), *prime)
	// c, _ := NewFieldElement(10, 13)

	d, _ := a.Mul(b)
	// res := c.IsEqual(&d)
	// fmt.Println(res)
	d.print()
}

// test field exponentation
func test2() {
	var a, _ = NewFieldElement(*big.NewInt(3), *big.NewInt(13))

	d, err := a.Pow(-3)
	if err != nil {
		fmt.Println(err)
	}

	e := a.IsEqual(&d)
	fmt.Println(e)
}

// test point exist on curve
func test3() {
	prime := big.NewInt(223)
	a, _ := NewFieldElement(*big.NewInt(5), *prime)
	b, _ := NewFieldElement(*big.NewInt(7), *prime)
	x1, er := NewFieldElement(*big.NewInt(5), *prime)
	if er != nil {
		fmt.Println(er)
		return
	}
	y1, er := NewFieldElement(*big.NewInt(7), *prime)
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
