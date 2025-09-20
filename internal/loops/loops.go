package loops

import "fmt"

func RunLoops() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)
	}

	ix := 0
	for ix < 5 {
		fmt.Printf("%d ", ix)
		ix++
	}
	fmt.Println()
}

func RunRangeLoops() {
	numbers := []int{11, 22, 33, 44, 55}
	for i, number := range numbers {
		fmt.Printf("%d: %d\n", i, number)
	}
}
