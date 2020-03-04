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
		fmt.Println("You have chosen " + choice)
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

// PerformFieldElement test module for finite fields
func PerformFieldElement() {
	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("Finite Fields\n\n")
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
		fmt.Println("You have chosen " + choice)
	case "c":
		fmt.Println("You have chosen " + choice)
	case "d":
		fmt.Println("You have chosen " + choice)
	case "e":
		fmt.Println("You have chosen " + choice)
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

//AddFe test module for finite field addition
func AddFe() {
	ClearScreen()
	readerFe := bufio.NewReader(os.Stdin)

	fmt.Print("Finite Field Addition\n\n")

	fmt.Print("Enter field size: ")

	sizeIn, _ := readerFe.ReadString('\n')
	sizeIn = strings.TrimSuffix(sizeIn, "\n")
	// size, _ := strconv.ParseFloat(sizeIn, 64)
	fmt.Print("\nField size is: " + sizeIn + "\n\n")

	fmt.Print("Enter first number: ")
	num1, _ := readerFe.ReadString('\n')
	num1 = strings.TrimSuffix(num1, "\n")
	n1, _ := strconv.ParseFloat(num1, 64)
	elem1, err := NewFieldElement(n1, 5)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(elem1.print())

	fmt.Print("Enter another number: ")
	num2, _ := readerFe.ReadString('\n')
	num2 = strings.TrimSuffix(num2, "\n")
	n2, _ := strconv.ParseFloat(num2, 64)
	elem2, err := NewFieldElement(n2, 5)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(elem2.print())

	res, err := elem1.Add(elem2)
	if err != nil {
		fmt.Println(err)
	}
	s := fmt.Sprintf("%.f", res)
	fmt.Print("\n" + num1 + "+" + num2 + " = " + s + "\n\n\n")

	defer PerformFieldElement()
}
