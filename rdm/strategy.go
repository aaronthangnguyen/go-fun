// https://github.com/tmrts/go-patterns/blob/master/behavioral/strategy.md
package main

import "fmt"

func main() {
	add := Operation{Operator: Addition{}}
	fmt.Printf("%d + %d = %d\n", 3, 5, add.Operate(3, 5))

	mul := Operation{Operator: Multiplication{}}
	fmt.Printf("%d * %d = %d\n", 3, 5, mul.Operate(3, 5))
}

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(leftValue, rightValue int) int {
	return leftValue + rightValue
}

type Multiplication struct{}

func (Multiplication) Apply(leftValue, rightValue int) int {
	return leftValue * rightValue
}
