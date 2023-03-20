// https://hackernoon.com/go-design-patterns-an-introduction-to-builder
package main

import "fmt"

func main() {
	rb := NewBuilder().
		Email("foo@bar.com").
		Balance(17.99).
		City("Orlando").
		Build()
	fmt.Printf("%+v\n", rb)
}

type Response struct {
	User    User
	Account Account
	Address Address
}

type User struct {
	Email string
}

type Account struct {
	Balance float64
}

type Address struct {
	City string
}

type ResponseBuilder struct {
	Response Response
}

func NewBuilder() *ResponseBuilder {
	return &ResponseBuilder{
		Response: Response{},
	}
}

func (rb *ResponseBuilder) Email(email string) *ResponseBuilder {
	rb.Response.User.Email = email
	return rb
}

func (rb *ResponseBuilder) Balance(balance float64) *ResponseBuilder {
	rb.Response.Account.Balance = balance
	return rb
}

func (rb *ResponseBuilder) City(city string) *ResponseBuilder {
	rb.Response.Address.City = city
	return rb
}

func (rb *ResponseBuilder) Build() Response {
	return rb.Response
}
