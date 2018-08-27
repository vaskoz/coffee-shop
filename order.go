package main

import (
	"time"
)

// PUBLIC INTERFACE
//

// Order represents a customer order
// STARTIFACE OMIT
type Order interface {
	//unique()
}

// STOPIFACE OMIT

// HIDDEN STRUCT IMPLEMENTING INTERFACE
//

// STARTSTRUCT OMIT
type order struct {
	taype string
	time  time.Duration
}

// STOPSTRUCT OMIT

// METHODS
//

//func (o *order) unique() {}

// PACKAGE FUNCTIONS
//

// NewOrder creates a new order to be filled
// STARTNEW OMIT
func NewOrder(taype string, time time.Duration) Order {
	return &order{taype: taype, time: time}
}

// STOPNEW OMIT
