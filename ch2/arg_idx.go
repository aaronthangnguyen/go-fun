// 2.70
package main

import "fmt"

func main() {
	fmt.Printf("in order: %s %s %s %s\n", "one", "two", "three", "four")
	fmt.Printf("explicit: %[4]s %[3]s %[2]s %[0]s\n", "one", "two", "three", "four")
}
