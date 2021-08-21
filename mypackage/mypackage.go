package mypackage

import "fmt"

//CarPublic
type CarPublic struct {
	Brand string
	Year int
}

//PrintMessage: Print message in console
func PrintMessage(message string) {
	fmt.Println(message)
}