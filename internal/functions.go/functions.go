package functions

import (
	"fmt",
	"strings"
)

func SayGreeting(n string) {
	fmt.Println("Hello", n)
}

func SayToAll(n []string, f func(string)) {
	for _, name := range n {
		f(name)
	}
}

func GetInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, name := range names {
		initials = append(initials, string(name[0]))
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	
	return initials[0], ""
}
