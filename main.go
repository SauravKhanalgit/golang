package main

import "fmt"

var outside = "you can't use := outside function"

// Dingdong := "hello"

func main() {

	// strings
	var nameOne string = "mario"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "Browser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi"                                //At terminal
	fmt.Println(nameOne, nameTwo, nameThree, nameFour) //go run main.go

	//int
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 30

	fmt.Println(numOne, numTwo, numThree)

	// float
	var scoreOne float32 = 23323.3
	var scoreTwo float64 = 34534534.54

	scoreThree := 342343.45

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
