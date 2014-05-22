package store

import (
	"errors"
	"github.com/vaskoz/coffee-shop/barista"
	"github.com/vaskoz/coffee-shop/customer"
	"github.com/vaskoz/coffee-shop/order"
	"time"
)

// INTERFACES
//

type Store interface {
	Open(chan<- string) error
	CloseAfter(time time.Duration)
	Customers([]customer.Customer)
	Baristas([]barista.Barista)
	isUnique()
}

// HIDDEN STRUCTS THAT IMPLEMENT INTERFACES
//

type store struct {
	baristas  []barista.Barista
	customers []customer.Customer
	openTime  time.Duration
	open      func(*store) error
  orderQueue chan order.Order
  baristaQueue chan barista.Barista
  customerQueue chan customer.Customer
  outputChan chan<- string
}

// HIDDEN DEFAULT METHOD IMPLEMENTATIONS
//

var defaultOpen = func(s *store) error {
	if s.openTime == 0 || s.customers == nil || s.baristas == nil {
		return errors.New("Can't open store. Not configured properly")
	}
  timeout := time.After(s.openTime)
  for {
    select {
    case <-timeout:
      return nil
    case c := <-s.customerQueue:
      b := <-s.baristaQueue
      go func() {
        o := b.MakeOrder(c.PlaceOrder())
        s.baristaQueue <- b
        s.outputChan <- c.EnjoyBeverage(o, b)
        s.customerQueue <- c
      }()
    }
  }
}

// METHODS
//

func (s *store) isUnique() {}

func (s *store) Open(outputChan chan<- string) error {
  s.outputChan = outputChan
	return s.open(s)
}

func (s *store) CloseAfter(time time.Duration) {
	s.openTime = time
}

func (s *store) Customers(customers []customer.Customer) {
	s.customers = customers
  s.customerQueue = make(chan customer.Customer, len(customers) + 1)
  for _, c := range s.customers {
    s.customerQueue <- c
  }
}

func (s *store) Baristas(baristas []barista.Barista) {
	s.baristas = baristas
  s.baristaQueue = make(chan barista.Barista, len(baristas) + 1)
  for _, b := range s.baristas {
    s.baristaQueue <- b
  }
}

// CONSTRUCTOR/FACTORY/BUILDER
// OTHER PACKAGE LEVEL EXPORTED FUNCTIONS

func New() Store {
	return &store{open: defaultOpen}
}
