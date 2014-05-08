package main

import (
	"fmt"
	"github.com/vaskoz/coffee-shop/barista"
	"github.com/vaskoz/coffee-shop/customer"
	"github.com/vaskoz/coffee-shop/store"
	"time"
)

// Purely for CLI purposes, so func main has something to start
var startCLI = func() {
	s := store.New()
	s.CloseAfter(10 * time.Second)
	s.Customers(customer.RandomGroupOf(20))
	s.Baristas(barista.RandomGroupOf(10))
	// Open() will block for the amount of time specificed by CloseAfter()
	// Unless there was an error
	err := s.Open()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	startCLI()
}
