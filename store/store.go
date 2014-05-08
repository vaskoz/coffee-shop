package store

import (
	"errors"
	"github.com/vaskoz/coffee-shop/barista"
	"github.com/vaskoz/coffee-shop/customer"
	"time"
)

// INTERFACES
//

type Store interface {
	Open() error
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
}

// HIDDEN DEFAULT METHOD IMPLEMENTATIONS
//

var defaultOpen = func(s *store) error {
	if s.openTime == 0 || s.customers == nil || s.baristas == nil {
		return errors.New("Can't open store. Not configured properly")
	}
	return nil
}

// METHODS
//

func (s *store) isUnique() {}

func (s *store) Open() error {
	return s.open(s)
}

func (s *store) CloseAfter(time time.Duration) {
	s.openTime = time
}

func (s *store) Customers(customers []customer.Customer) {
	s.customers = customers
}

func (s *store) Baristas(baristas []barista.Barista) {
	s.baristas = baristas
}

// CONSTRUCTOR/FACTORY
// OTHER PACKAGE LEVEL EXPORTED FUNCTIONS

func New() Store {
	return &store{open: defaultOpen}
}
