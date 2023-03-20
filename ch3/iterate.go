// 3.38

package main

import "fmt"

func main() {
	names := []string{"Kurt", "Janis", "Jimi", "Amy"}
	for i, name := range names[1:] {
		fmt.Println(i, name)
	}
}
