// 3.15
package main

import "fmt"

func main() {
	var names []string
	names = append(names, "Kris")
	fmt.Println(names)

	more := []string{"Janis", "Jimi"}

	for _, name := range more {
		names = append(names, name)
	}

	fmt.Println(names)
}
