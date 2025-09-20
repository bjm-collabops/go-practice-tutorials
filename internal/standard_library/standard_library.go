package standard_library

import (
	"fmt"
	"sort"
	"strings"
)

func RunStringStandardLibrary() {
	greeting := "Hello, World!"
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "World"))
	fmt.Println(strings.Split(greeting, " "))
	fmt.Println(strings.Join([]string{"Hello", "World"}, " "))

	fmt.Println(greeting + "-> shows original string")

	fmt.Println(sort.StringsAreSorted([]string{"a", "b", "c"}))
	fmt.Println(sort.StringsAreSorted([]string{"a", "d", "c"}))
}
