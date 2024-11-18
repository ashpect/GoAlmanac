package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("go" + "lang")
	// Integers and floats.
		fmt.Println("1+1 =", 1+1)
		fmt.Println("7.0/3.0 =", 7.0/3.0)
	// Booleans, with boolean operators as you’d expect.
		fmt.Println(true && false)
		fmt.Println(true || false)
		fmt.Println(!true)

	// VARIABLES : 
	// Variables declared without a corresponding initialization are zero-valued.
	// For example, zero value for map is nil.
	var a string = "initial"
	var b, c int = 1, 2
	d,e := 10/3,2
	fmt.Println(a,b,c,d,e)

	// CONSTS :
	// Declares a constant value. A const statement can appear anywhere a var statement can.
	// Constant expressions perform arithmetic with arbitrary precision.
	// To have proper precision, use big package. eg :
	const n = 500000000
	const m = 3e20 / n
	fmt.Println(m)
	fmt.Println(int64(m))


	// FOR LOOPS :
	i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j < 3; j++ {
        fmt.Println(j)
    }

    for i := range 3 {
        fmt.Println("range", i)
    }

    for {
        fmt.Println("loop")
        break
    }

    for n := range 6 {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }

	// IF STATEMENTS :
	// No ternary "?" operator exists in Go.
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	
	// STRINGS :
	// Unicode package gives fucntions to work with runes.
	// IsDigit, IsLetter, IsLower, IsUpper, ToUpper, ToLower, ToTitle


}