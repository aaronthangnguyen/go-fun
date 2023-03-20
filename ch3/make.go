// 3.25
package main

import "fmt"

func main() {
	a := make([]string, 1, 3)
	// slice with length of 1 and capacity of 3
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
