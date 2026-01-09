package main

import "fmt"

func main() {
	age := 20
	name := "Saurav"
	//print
	fmt.Print("Hello, ")

	fmt.Print("World! \n")
	fmt.Print("new line\n")
	// println
	fmt.Println("Hello guys")
	fmt.Println("Bye guys")
	fmt.Println("My age is", age, "and my name is", name)

	// printf (formatted strings) %_ = format specifier
	fmt.Printf("My name is %v and my age is %v\n", name, age)
	// %q for quote, but for integer it doesn't work
	fmt.Printf("My name is %q and my age is %q\n", name, age)
	fmt.Printf("Age is of type %T\n", age)

	//float
	fmt.Printf("The score is %0.2f \n", 255.22)

	//sprintf save formatted string
	var str = fmt.Sprintf("My name is %v and my age is %v\n", name, age)
	fmt.Println("The saved string is: ", str)

}
