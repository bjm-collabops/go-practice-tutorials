package main

import (
	"go-practice/internal/examples"
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
	examples.DemonstrateBasicPrinting()
	examples.DemonstrateFormattedPrinting()

	// Demonstrate arrays and slices
	types.DemonstrateArrays()

}
