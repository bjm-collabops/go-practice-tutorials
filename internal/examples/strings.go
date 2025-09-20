package examples

import "fmt"

// DemonstrateBasicPrinting shows basic printing functions
func DemonstrateBasicPrinting() {
	fmt.Print("hello", "world", "\n")
}

// DemonstrateFormattedPrinting shows formatted printing with Printf and Sprintf
func DemonstrateFormattedPrinting() {
	myName := "John"
	myAge := "30"
	// println doens't have printf formatting
	// printf has a special "Quote format" which allows you to print strings with quotes
	fmt.Printf("My name is %q and my age is %q\n", myName, myAge)
	// You can save your formatted string to a variable
	str := fmt.Sprintf("My name is %q and my age is %q\n", myName, myAge)
	fmt.Println(str)
}
