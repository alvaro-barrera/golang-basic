package main

import "fmt"

func main() {
	//const
	const pi float64 = 3.14
	const pi_2 = 3.14

	fmt.Println("Hello world")
	fmt.Println("pi: ",pi)
	fmt.Println("pi_2: ",pi_2)

	//int
	base := 12
	var height int = 14
	var area int 

	fmt.Println(base,height,area)

	//zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a,b,c,d)

	//area of a square
	const base_square = 15
	area_square := base_square * base_square
	println("Area: ", area_square)

	x := 30
	y := 5

	//sum
	result := x + y
	println("Sum: ",result)

	//subtraction
	result = x - y
	println("Subtraction: ",result)

	//multiplication
	result = x * y
	println("Multiplication: ",result)

	//division
	result = x / y
	println("Division: ",result)

	//remainder
	result = x % y
	println("Remainder: ",result)

	//incremental
	x++
	println("Incremental: ",x)

	//decremental
	x--
	println("Decremental: ",x)
}