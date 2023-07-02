package main

import (
	"fmt"
	"runtime"
)

// for loop
func forLoop() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("This sum is: %d", sum)
}

// for is Go like `while` in another language
func forLoopWithoutInitAndPostStatement() {
	sum := 0
	for sum < 100 {
		sum += 10
	}

	fmt.Printf("\nThis sum without init and post statement is: %d", sum)
}

// infinity loop. Don't run it if you want machine is crash.
func inifinityLoop() {
	for {
	}
}

// conditional in Go
func ifConditional(x int) {
	if x < 0 {
		fmt.Println("\nThe number is negative.")
	}
}

func ifConditionalWithShortStatement(x, n int) {
	if v := x % n; v == 0 { // v is only scope until the end of the `if`
		fmt.Printf("%d divisible by %d\n", x, n)
	}
}

func ifElseConditional(x, n int) {
	if v := x % n; v == 0 { // v is only scope until the end of the `if`
		fmt.Printf("%d divisible by %d\n", x, n)
	} else {
		fmt.Printf("%d not divisible by %d, value is %d", x, n, v)
	}
}

// Switch
// Go only run selected case. In short, `break` is auto provided in Go.
func switchConditional() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("\n%s. \n", os)
	}
}

// case 1 not execute
func anotherSwitchCase() {
	switch v := 0; v {
	case 0:
	case 1:
		fmt.Println("Value is 1.")
	default:
		fmt.Println("Value is default")
	}
}

// Switch with multi case
func switchWithMultiCase(day string) {
	switch day {
	case "today", "tomorrow":
		fmt.Println("today or tomorrow.")
	case "yesterday":
		fmt.Println("yesterday")
	default:
		fmt.Println("unknown.")
	}
}

// Switch without conditional like `switch true`

// `defer` in Go
// Defer in Go is awesome. It like await in Javascript but it await all not defer is called then it is called.
func deferAwesome() {
	var x byte = 10
	defer fmt.Printf("Value of x is: %d\n", x) // defer call's arguments are evaluated immediately, but the call is not executed
	x += 10
	fmt.Printf("Run first and defer run after. So value is: %d\n", x)
}

// So if have more defer? Stack defers is solution.
// defers is called in `last in first out`

func main() {
	forLoop()
	forLoopWithoutInitAndPostStatement()
	ifConditional(-4)
	ifConditionalWithShortStatement(10, 5)
	ifElseConditional(12, 5)
	switchConditional()
	anotherSwitchCase()
	switchWithMultiCase("tomorrow")
	deferAwesome()
}
