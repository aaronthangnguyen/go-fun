// 2.66
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "Kurt",
		Age:  27,
	}

	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
}
