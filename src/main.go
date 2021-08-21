package main

import (
	"Code/golang_first_app/mypackage"
	"fmt")


func main() {
	var myCarPublic mypackage.CarPublic
	myCarPublic.Brand = "Ferrari"
	myCarPublic.Year = 2021
	fmt.Println(myCarPublic)
	mypackage.PrintMessage("Hello World")
}