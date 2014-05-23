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
	outputChan := make(chan string, 100)
	s := store.New()
	s.CloseAfter(20 * time.Second)
	s.Customers(customer.RandomGroupOf(20))
	s.Baristas(barista.RandomGroupOf(10))
	go func() {
		err := s.Open(outputChan)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("*** Store is closed now ***")
		}
	}()
	for {
		select {
		case str := <-outputChan:
			fmt.Println(str)
		case <-time.After(5 * time.Second):
			return
		}
	}
}

func main() {
	startCLI()
}
