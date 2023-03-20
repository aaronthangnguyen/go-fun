// 2.22
package main

import "fmt"

func main() {
	a := "one 🐘 and three 🐋"
	for i, c := range a {
		fmt.Printf("%d %s\n", i, string(c))
        
	}
}
