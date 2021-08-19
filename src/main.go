package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func doubleValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
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
		fmt.Println("Sin condición")
	}
}