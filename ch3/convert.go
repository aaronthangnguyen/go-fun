// 3.31

package main

import "fmt"

func slicesOnly(names []string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	names := [4]string{"Kurt", "Janis", "Jimi", "Amy"}

	// not working
	// slicesOnly(names)
	// slicesOnly([]string(names))

	slicesOnly(names[:])
}
