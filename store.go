package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// INTERFACES
//

// Store represents a place where customers, baristas and orders are filled
type Store interface {
	Open(context.Context) (<-chan struct{}, error)
	CloseAfter(time time.Duration)
	Customers([]Customer)
	Baristas([]Barista)
	//isUnique()
}

// HIDDEN STRUCTS THAT IMPLEMENT INTERFACES
//

type store struct {
	baristas      []Barista
	customers     []Customer
	openTime      time.Duration
	baristaQueue  chan Barista
	customerQueue chan Customer
	logger        *log.Logger
}

// HIDDEN DEFAULT METHOD IMPLEMENTATIONS
//

// METHODS
//

//func (s *store) isUnique() {}

func (s *store) Open(ctx context.Context) (<-chan struct{}, error) {
	if s.openTime == 0 || len(s.customers) == 0 || len(s.baristas) == 0 {
		return nil, errors.New("can't open store. not configured properly")
	}

	ctx, cancel := context.WithTimeout(ctx, s.openTime)
	done := make(chan struct{}, 1)

	var wg sync.WaitGroup

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.logger.Println("Store is closing")
				wg.Wait()
				cancel()
				done <- struct{}{}

				return
			case c := <-s.customerQueue:
				b := <-s.baristaQueue

				wg.Add(1)

				go func() {
					o := b.MakeOrder(c.PlaceOrder())
					s.baristaQueue <- b
					s.logger.Println(c.EnjoyBeverage(o, b))
					s.customerQueue <- c

					wg.Done()
				}()
			}
		}
	}()

	return done, nil
}

func (s *store) CloseAfter(time time.Duration) {
	s.openTime = time
}

func (s *store) Customers(customers []Customer) {
	s.customers = customers
	s.customerQueue = make(chan Customer, len(customers)+1)

	for _, c := range s.customers {
		s.customerQueue <- c
	}
}

func (s *store) Baristas(baristas []Barista) {
	s.baristas = baristas
	s.baristaQueue = make(chan Barista, len(baristas)+1)

	for _, b := range s.baristas {
		s.baristaQueue <- b
	}
}

// CONSTRUCTOR/FACTORY/BUILDER
// OTHER PACKAGE LEVEL EXPORTED FUNCTIONS

// NewStore creates a new store
// Pass it a logger
func NewStore(logger *log.Logger) Store {
	return &store{logger: logger}
}
