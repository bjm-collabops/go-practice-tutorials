package types

import "fmt"

// DemonstrateVariableDeclarations shows different ways to declare variables
func DemonstrateVariableDeclarations() {
	// strings (explicit and implicit typed)
	var nameOne string = "hello"
	var nameThree string
	fmt.Println(nameOne, nameThree)

	nameOne = "changed"
	nameThree = "changed2"
	fmt.Println(nameOne, nameThree)

	// short hand for declaring and initializing a variable
	nameFour := "hello4"
	fmt.Println(nameFour)
}

// DemonstrateIntegerTypes shows different integer types and their ranges
func DemonstrateIntegerTypes() {
	// ints
	var ageOne int = 40
	var ageTwo = 50
	ageThree := 60

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory (var could be useful here)
	var numOne int8 = 25     // 8 bits (-128 to 127)
	var numTwo int16 = 256   // 16 bits (-32768 to 32767)
	var numThree int32 = 256 // 32 bits (-2147483648 to 2147483647)
	// default int is int64
	var numFour int64 = 256 // 64 bits (-9223372036854775808 to 9223372036854775807)
	fmt.Println(numOne, numTwo, numThree, numFour)

	var numFive uint8 = 255   // 8 bits (0 to 255)
	var numSix uint16 = 256   // 16 bits (0 to 65535)
	var numSeven uint32 = 256 // 32 bits (0 to 4294967295)
	var numEight uint64 = 256 // 64 bits (0 to 18446744073709551615)
	numFive += 1              // overflows to 0
	fmt.Println(numFive, numSix, numSeven, numEight)
}

// DemonstrateFloatTypes shows float types
func DemonstrateFloatTypes() {
	var scoreOne float32 = 25.98  // 32 bits
	var scoreTwo float64 = 256.98 // 64 bits (default float)
	fmt.Println(scoreOne, scoreTwo)
}

// DemonstrateRuneType shows rune type (alias for int32)
func DemonstrateRuneType() {
	var scoreThree rune = 'a'       // 32 bits (rune is an alias for int32)
	fmt.Println(scoreThree)         // prints the ascii value of a
	fmt.Println(string(scoreThree)) // prints the character a
}
