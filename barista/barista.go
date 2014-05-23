package barista

import (
	"fmt"
	"github.com/vaskoz/coffee-shop/order"
	"time"
)

type Barista interface {
	MakeOrder(order.Order) order.Order
	Id() int
	String() string
}

type barista struct {
	speed        int
	processOrder func(order.Order)
	id           int
}

func (b *barista) MakeOrder(order order.Order) order.Order {
	time.Sleep(time.Duration(b.id) * time.Second)
	return order
}

func (b *barista) Id() int {
	return b.id
}

func (b *barista) String() string {
	return fmt.Sprintf("Barista %d", b.id)
}

var New = func(id int) Barista {
	return &barista{id: id, speed: 2}
}

var RandomGroupOf = func(n int) []Barista {
	baristas := make([]Barista, n)
	for i := 0; i < n; i++ {
		baristas[i] = New(i + 1)
	}
	return baristas
}
