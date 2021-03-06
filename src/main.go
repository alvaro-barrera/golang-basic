package main

import (
	"fmt"
	"strconv"
)


func normalFunction(message string) {
	fmt.Println(message)
}

func doubleValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func isPalindromo(text string) {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		normalFunction("Es palindromo")
	}else{
		normalFunction("No es palindromo")
	}
}

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

	normalFunction("Message")

	value := doubleValue(6)

	fmt.Println("Value: ",value)

	value1, _ := doubleReturn(4)

	fmt.Println("Value1: ",value1)

	//For
	for i := 0; i < 12; i++ {
		fmt.Println(i)
	}

	//For while
	count := 0
	for count < 15 {
		fmt.Println(count)
		count++
	}

	//For forever
	// countForever :=0
	// for {
	// 	fmt.Println(countForever)
	// 	countForever++		
	// }

	//If
	valueIf1 := 1
	valueIf2 := 2
	if valueIf1 == 1 {
		fmt.Println("Es 1")
	}else{
		fmt.Println("No es 1")
	}
	if valueIf1 == 1 && valueIf2 == 3{
		fmt.Println("Es verdad, AND")
	}else{
		fmt.Println("Es falso, AND")
	}

	if valueIf1 == 0 || valueIf2 == 2{
		fmt.Println("Es verdad, OR")
	}else{
		fmt.Println("Es falso, OR")
	}

	modulo := 7 % 2
	switch modulo {
	case 0:
		normalFunction("Es par")
	default:
		normalFunction("Es impar")
	}

	valueSwitch := 60
	switch {
	case valueSwitch > 100:
		fmt.Println("Es mayor a 100")
	case valueSwitch <0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("Sin condici??n")
	}

	//Defer
	defer fmt.Println("Close")
	fmt.Println("After defer")

	//Continue - Break
	for i := 0; i < 10; i++ {
		if i == 4 {
			normalFunction("Continue")
			continue
		}

		if i == 7 {
			normalFunction("Break")
			break
		}
	}

	//Array
	var array [5]int
	array[0] = 3
	array[1] = 1
	fmt.Println(array, len(array), cap(array))

	//Slice
	slice := []int{0,1,2,3}
	fmt.Println(slice, len(slice), cap(slice))

	//Slice methods
	normalFunction(strconv.Itoa(slice[0]))
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[3:])

	//Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//Append new list
	newSlice := []int{8,9,10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	//Range
	sliceRange := []string{"Hola", "c??mo", "est??s?"}

	for i := range sliceRange {
		fmt.Println(i)
	}

	//Palindromo
	isPalindromo("amor a roma")

	//map
	array_ := make(map[string]int)

	array_["Jose"] = 2
	array_["Juan"] = 2

	fmt.Println(array_)

	for i, v := range array_ {
		fmt.Println(i, v)
	}

	value_array, ok := array_["Jose"]
	fmt.Println(value_array, ok)
}