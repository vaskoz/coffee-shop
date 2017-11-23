package main

import (
	"fmt"
)

// Customer consumes coffee
type Customer interface {
	PlaceOrder() Order
	EnjoyBeverage(Order, interface{}) string
}

type customer struct {
	id int
}

// PlaceOrder lets a customer create an order for a Barista
func (c *customer) PlaceOrder() Order {
	return NewOrder("latte", 2)
}

// EnjoyBeverage lets a customer enjoy a fulfilled order
func (c *customer) EnjoyBeverage(order Order, provider interface{}) string {
	return fmt.Sprintf("Customer %d says Yum and thanks to %v", c.id, provider)
}

// NewCustomer creates a new coffee buying customer
func NewCustomer(id int) Customer {
	return &customer{id: id}
}

// RandomGroupOfCustomers creates a group of customers
func RandomGroupOfCustomers(n int) []Customer {
	customers := make([]Customer, n)
	for i := 0; i < n; i++ {
		customers[i] = NewCustomer(i + 1)
	}
	return customers
}
