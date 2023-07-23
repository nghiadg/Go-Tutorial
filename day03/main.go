package main

import "fmt"

func pointer() {
	var i int = 42
	var p *int = &i
	fmt.Println(p) // It's return a memory address of i variable. Example: 0xc000016088
	fmt.Println(*p)
	i = 30
	fmt.Println(*p)
}

func structs() {
	// Struct like class in OOP or type of Object in Typescript.
	type Vertex struct {
		X int
		Y int
	}

	var vertex Vertex = Vertex{1, 2}
	fmt.Println(vertex)
	fmt.Println(vertex.X)
	fmt.Println(vertex.Y)
	vertex.X = 4
	fmt.Println(vertex.X)
	var p *Vertex = &vertex
	p.X = 10
	fmt.Println(p.X)
}

func array() {
	var arr [3]string // Array length is part of it's type. So arrays cannot be resized.
	arr[0] = "Hello"
	arr[1] = "World"

	fmt.Println(arr[2] == "")

	// You can use this syntax to defined a array and value of it
	primes := [3]int{1, 2, 3}
	// Not using literal
	var anotherPrimes [3]int = [3]int{1, 2, 3}
	fmt.Println(anotherPrimes)
	fmt.Println(primes)

}

func slices() {
	// Slice is dynamically sized.
	var sl []int = []int{1, 2, 3, 4}
	fmt.Println(sl)
	/* You can slice a array like Javascript slice work by syntax. arr[low, hight].
	Note: it don't contain `hight` value, let try example below
	*/
	arr := [5]int{1, 2, 3, 4, 5}
	var slArr []int = arr[1:4] // -> [2,3,4] not contain `5`
	fmt.Println(slArr)

	//TODO: Explain deep dive into how array and slice is stored in memory
}

func main() {
	// structs()
	// array()
	slices()
}
