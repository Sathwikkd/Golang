package main

import (
	"fmt"
)

// function calls
func main() {
	basicVariable()
	declarationInit()
	inference()
	shorthanded()
	multiVar1()
	multiVar2()
	zeroVal()
}

// basic variable declaration
func basicVariable() {
	var number int
	number = 10
	fmt.Println(number)
}

// declaration with initialization
func declarationInit() {
	var name string = "string"
	println(name)
}

// type inference
func inference() {
	var weight = 20.76
	fmt.Println(weight)
}

// shorthanded
func shorthanded() {
	isTrue := true
	floatval := 50.55
	integer := 50
	strings := "hello Go.."
	fmt.Println(strings)
	fmt.Println(floatval)
	fmt.Println(integer)
	fmt.Println(isTrue)
}

// multi variable declaration type-1
func multiVar1() {
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)
}

// multi variable declaration type-2
func multiVar2() {
	x, y, z := 10, "hello", true
	fmt.Println(x, y, z)
}

// zero values
func zeroVal() {
	var i int
	var j float32
	var k string
	var l bool
	fmt.Println(" int is", i, "\n", "float is", j, "\n", "string is", k, "\n", "bool is", l)
}
