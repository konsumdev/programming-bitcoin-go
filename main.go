package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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

// this is the main
func main() {
	defer prompt()
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
		defer PerformFieldElement()
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
	fmt.Println("Point 1")
	x, y, a, b := GetCoors()
	point, err := NewPoint(x, y, a, b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Point 2")
	x2, y2, a2, b2 := GetCoors()

	// fmt.Printf("Point(%s, %s)_%s_%s\n\n", x, y, a, b)
	// fmt.Print(" + ")
	// fmt.Printf("Point(%s, %s)_%s_%s\n\n", x2, y2, a2, b2)

	point2, err := NewPoint(x2, y2, a2, b2)
	if err != nil {
		fmt.Println(err)
	}

	res, _ := point.Add(point2)
	res.print()
}

// PerformFieldElement test module for finite fields
func PerformFieldElement() {
	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("\nFinite Fields\n\n")
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
	s := fmt.Sprintf("%.f", res)
	fmt.Println("Result: " + s)

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
	s := fmt.Sprintf("%.f", res)
	fmt.Println("Result: " + s)

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
	s := fmt.Sprintf("%.f", res)
	fmt.Println("Result: " + s)

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
	s := fmt.Sprintf("%.f", res)
	fmt.Println("Result: " + s)

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
	s := fmt.Sprintf("%.f", res)
	fmt.Println("Result: " + s)
	// fmt.Print("\n" + num1 + "+" + num2 + " = " + s + "\n\n\n")

	defer PerformFieldElement()
}
