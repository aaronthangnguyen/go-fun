// 3.16
package main

import "fmt"

func main() {
	var names []string
	names = append(names, "Kris")
	fmt.Println(names)

	more := []string{"Janis", "Jimi"}

	names = append(names, more...)
	for _, name := range more {
		names = append(names, name)
	}
	for _, name := range more {
		names = append(names, name)
	}

	fmt.Println(names)
}
