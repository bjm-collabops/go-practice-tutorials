package types

import "fmt"

func DemonstrateBooleans() {
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)

	age := 20
	fmt.Println(age >= 18)

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}

	names := []string{"John", "Jane", "Jim", "Jill"}
	for i, name := range names {
		if i == 1 {
			fmt.Println("Skipping Jane")
			// continue will skip the rest of the loop and go to the next iteration
			continue
		}
		fmt.Printf("%d: %s\n", i, name)
	}
}
