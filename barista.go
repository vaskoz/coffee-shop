package main

import (
	"fmt"
	"time"
)

// Barista makes coffee
type Barista interface {
	MakeOrder(Order) Order
	String() string
}

type barista struct {
	speed int
	id    int
}

func (b *barista) MakeOrder(order Order) Order {
	time.Sleep(time.Duration(b.id) * time.Second)
	return order
}

func (b *barista) String() string {
	return fmt.Sprintf("Barista %d", b.id)
}

// NewBarista creates a new Barista
var NewBarista = func(id int) Barista {
	return &barista{id: id, speed: 15}
}

// RandomGroupOfBaristas creates a random group of baristas
var RandomGroupOfBaristas = func(n int) []Barista {
	baristas := make([]Barista, n)
	for i := 0; i < n; i++ {
		baristas[i] = NewBarista(i + 1)
	}
	return baristas
}
