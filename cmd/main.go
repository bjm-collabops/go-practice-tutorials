package main

import (
	"go-practice/internal/functions.go"
	"go-practice/internal/loops"
	"go-practice/internal/standard_library"
	"go-practice/internal/types"
)

// go run cmd/main.go
func main() {
	// Demonstrate variable declarations
	types.DemonstrateVariableDeclarations()
	types.DemonstrateIntegerTypes()
	types.DemonstrateFloatTypes()

	types.DemonstrateRuneType()

	// Demonstrate basic printing
	types.DemonstrateBasicPrinting()
	types.DemonstrateFormattedPrinting()

	// Demonstrate arrays and slices
	types.DemonstrateArrays()

	// Demonstrate string standard library
	standard_library.RunStringStandardLibrary()

	loops.RunLoops()
	loops.RunRangeLoops()

	functions.SayGreeting("John")
	functions.SayToAll([]string{"John", "Jane", "Jim", "Jill"}, functions.SayGreeting)
	functions.getInitials("John Doe")
	functions.getInitials("John")
	functions.getInitials("John Doe Smith")
}
