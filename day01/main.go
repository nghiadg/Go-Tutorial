package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

/*
`naked` return
Go's return values may be named.
A `return` statement without arguments return the named return values.
*/
func split(sum int) (a, b int) {
	a = sum / 5
	b = sum - a
	return
}

// Variables in Go
// To declare a variable in go use `var <name> type`.
var isChecked bool

func check() bool {
	return isChecked
}

// Go can infer type of initialized variable. Let see example:
var inferVar = true

// Can declare multi variables and type of them is last.
var python, javascript, golang string = "python", "javascript", "golang"

//short variable. Just inside function, not allow outside.
// shortVar := "This is short variable"

func printShortVar() {
	shortVar := "This is short variable"
	fmt.Println(shortVar)
}

func basicTypes() {
	var xInt8 int8 = 127          // 1 bit sign and 7 bit to store value
	var xInt32 int32 = 2147483647 // 1 bit sign and 31 bit to store value

	var xByte byte = 127        // alias for int8
	var xRune rune = 2147483647 // alias for int32

	fmt.Println(xInt8, xInt32, xByte, xRune)
}

// declare variable and not init for it
// 0 for numeric
// false for boolean
// "" for string
func declareVarWithoutInit() {
	var i int
	var f float32
	var b bool
	var s string
	fmt.Println(i, f, b, s)
}

// when declare a variable without type. The type of variable is referred from value in right hand side.
// when the right hand side of declaration is typed, the new variable is of the same type
// when the right hand side contains an untyped numeric constant, the new variable maybe an `int`, `float64` or `complex128`

func declareVarWithoutType() {
	var a int
	b := a // type of b is int
	fmt.Println(b)

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Println(i, f, g)
}

// constants are declare like variable but using `const` keyword
// constants cannot be declared using the `:=` syntax

func declareConstant() {
	const HOURS_OF_DAY int8 = 24
	fmt.Println(HOURS_OF_DAY)
}

func main() {
	fmt.Println(sum(2, 5))
	fmt.Println(swap("Hello", "World"))
	fmt.Println(split(100))
	printShortVar()
	basicTypes()
	declareVarWithoutInit()
	declareVarWithoutType()
	declareConstant()
}
