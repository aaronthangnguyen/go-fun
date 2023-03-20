// 2.33
package main

import "fmt"

func main() {
	i, f, b, s := Values()

	fmt.Printf("var i %T: %v\n", i, i)
	fmt.Printf("var f %T: %f\n", f, f)
	fmt.Printf("var b %T: %v\n", b, b)
	fmt.Printf("var s %T: %q\n", s, s)
}

func Values() (int, float64, bool, string) {
	return 42, 3.14, true, "hello world"
}
