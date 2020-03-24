package main

/**
TODO

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)



var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //mac example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ClearScreen clears the cmd screen
func ClearScreen() {

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Platform unsupported! I can't clear terminal screen :(")
	}
}

func prompt() {
	fmt.Print("\n\nProgramming bitcoin in Go\n\n")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("a] FieldElement")
	fmt.Println("b] Point Addition")
	fmt.Println("x] Clear screen")
	fmt.Print("z] exit \n\n")
	fmt.Print("Run test for: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")

Loopers:
	switch choice {
	case "a":
		ClearScreen()
		fmt.Print("\nYou have chosen " + choice + "\n\n")
		// defer PerformFieldElement()
	case "b":
		ClearScreen()
		fmt.Print("\nYou have chosen " + choice + "\n\n")
		PerformPointAddition()
	case "x":
		ClearScreen()
		defer prompt()
	case "z":
		fmt.Println("exit")
		break Loopers
	default:
		ClearScreen()
		fmt.Print("Unrecognised input\n")
		defer prompt()
	}
}

// GetCoors test module for point addition
func GetCoors() (string, string, string, string) {
	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("Enter value for a: ")
	a, _ := readerFe.ReadString('\n')
	a = strings.TrimSuffix(a, "\n")

	fmt.Print("Enter value for b: ")
	b, _ := readerFe.ReadString('\n')
	b = strings.TrimSuffix(b, "\n")

	fmt.Print("Enter value for x: ")
	x, _ := readerFe.ReadString('\n')
	x = strings.TrimSuffix(x, "\n")

	fmt.Print("Enter value for y: ")
	y, _ := readerFe.ReadString('\n')
	y = strings.TrimSuffix(y, "\n")

	return x, y, a, b
}

// PerformPointAddition test module for point addition
func PerformPointAddition() {

	fmt.Print("\nPoint Addition\n\n")

	var point Point
	var point2 Point
	var err error

Ask:
	for {
		fmt.Println("Point 1")
		x, y, a, b := GetCoors()
		point, err = NewPoint(x, y, a, b)
		if err != nil {
			fmt.Println(err)
		} else {
			break Ask
		}
	}

AskFurther:
	for {
		fmt.Println("Point 2")
		x2, y2, a2, b2 := GetCoors()

		point2, err = NewPoint(x2, y2, a2, b2)
		if err != nil {
			fmt.Println(err)
		} else {
			break AskFurther
		}
	}

	res, _ := point.Add(&point2)
	res.print()

	defer prompt()
}



// PerformFieldElement test module for finite fields
func PerformFieldElement() {
	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("\n\nFinite Fields\n\n")
	fmt.Println("a] Addition")
	fmt.Println("b] Subtraction")
	fmt.Println("c] Multiplication")
	fmt.Println("d] Exponent")
	fmt.Println("e] Division")
	fmt.Println("x] Clear screen")
	fmt.Print("z] back \n\n")
	fmt.Print("Run test for: ")

	choice, _ := readerFe.ReadString('\n')
	choice = strings.TrimSuffix(choice, "\n")

	switch choice {
	case "a":
		AddFe()
	case "b":
		SubFe()
	case "c":
		MulFe()
	case "d":
		ExpFe()
	case "e":
		DivFe()
	case "x":
		ClearScreen()
		defer PerformFieldElement()
	case "z":
		ClearScreen()
		defer prompt()
	default:
		ClearScreen()
		fmt.Print("Unrecognised input\n\n")
		defer PerformFieldElement()
	}
}

// DivFe test module for finite field multiplication
func DivFe() {
	ClearScreen()
	fmt.Print("Finite Field Division\n\n")

	elem1, elem2 := InpReader()

	res, err := elem1.Div(elem2)
	if err != nil {
		fmt.Println(err)
	}

	// print result
	res.print()

	defer PerformFieldElement()
}


// ExpFe test module for finite field exponentiation
func ExpFe() {
	ClearScreen()
	fmt.Print("Finite Field Exponentiation\n\n")

	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("Enter field size: ")

	sizeIn, _ := readerFe.ReadString('\n')
	sizeIn = strings.TrimSuffix(sizeIn, "\n")
	size, _ := strconv.ParseFloat(sizeIn, 64)

	fmt.Print("\nField size is: " + sizeIn + "\n\n")
	var elem1 FieldElement
	var err error
Ask:
	for {
		fmt.Print("Enter number: ")
		num1, _ := readerFe.ReadString('\n')
		num1 = strings.TrimSuffix(num1, "\n")
		n1, _ := strconv.ParseFloat(num1, 64)
		elem1, err = NewFieldElement(n1, 5)
		if err != nil {
			fmt.Println(err)
		} else {
			break Ask
		}
	}

	res, err := elem1.Pow(size)
	if err != nil {
		fmt.Println(err)
	}
	// print result
	res.print()

	defer PerformFieldElement()
}

// MulFe test module for finite field multiplication
func MulFe() {
	ClearScreen()
	fmt.Print("Finite Field Multiplication\n\n")

	elem1, elem2 := InpReader()

	res, err := elem1.Mul(elem2)
	if err != nil {
		fmt.Println(err)
	}

	// print result
	res.print()

	defer PerformFieldElement()
}

// SubFe test module for finite field subtraction
func SubFe() {
	ClearScreen()
	fmt.Print("Finite Field Subtraction\n\n")

	elem1, elem2 := InpReader()

	res, err := elem1.Sub(elem2)
	if err != nil {
		fmt.Println(err)
	}

	// print result
	res.print()

	defer PerformFieldElement()
}

// InpReader test module for finite field addition
func InpReader() (f FieldElement, fe FieldElement) {

	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("Enter field size: ")

	sizeIn, _ := readerFe.ReadString('\n')
	sizeIn = strings.TrimSuffix(sizeIn, "\n")
	size, _ := strconv.ParseFloat(sizeIn, 64)

	fmt.Print("\nField size is: " + sizeIn + "\n")

	var elem1 FieldElement
	var elem2 FieldElement
	var err error

Ask:
	for {
		fmt.Print("Enter first number: ")
		num1, _ := readerFe.ReadString('\n')
		num1 = strings.TrimSuffix(num1, "\n")
		n1, _ := strconv.ParseFloat(num1, 64)
		elem1, err = NewFieldElement(n1, size)
		if err != nil {
			fmt.Println(err)
		} else {
			break Ask
		}
	}

AskTwice:
	for {
		fmt.Print("Enter another number: ")
		num2, _ := readerFe.ReadString('\n')
		num2 = strings.TrimSuffix(num2, "\n")
		n2, _ := strconv.ParseFloat(num2, 64)
		elem2, err = NewFieldElement(n2, size)
		if err != nil {
			fmt.Println(err)
		} else {
			break AskTwice
		}
	}

	return elem1, elem2
}

// AddFe test module for finite field addition
func AddFe() {
	ClearScreen()
	fmt.Print("Finite Field Addition\n\n")

	elem1, elem2 := InpReader()

	res, err := elem1.Add(elem2)
	if err != nil {
		fmt.Println(err)
	}

	// print result
	res.print()

	defer PerformFieldElement()
}


************


func testVerify() {
	z := hexToBigInt("0xbc62d4b80d9e36da29c16c5d4d9f11731f36052c72401a76c23c0fb5a9b74423")
	r := hexToBigInt("0x37206a0610995c58074999cb9767b87af4c4978db68c06e8e6e81d282047a7c6")
	s := hexToBigInt("0x8ca63759c1157ebeaec0d03cecca119fc9a75bf8e6d0fa65c841c8e2738cdaec")
	px := hexToBigInt("0x79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798")
	py := hexToBigInt("0x483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8")
	// sig := NewSignature(NewS256Field(*r), NewS256Field(*s))

	point := NewS256Point(*px, *py)

	// fmt.Println(new(big.Int).Exp(s, z, nil))

	ver := point.verify(z, s, r)

	fmt.Println(ver)
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
	b, _ := a.Pow(*big.NewInt(3))
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

	d, err := a.Pow(*big.NewInt(-3))
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
**/
