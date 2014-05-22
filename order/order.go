package order

import (
  "time"
)

// PUBLIC INTERFACE
//

type Callback func(Order) error

// STARTIFACE OMIT
type Order interface {
	Type() string
	Time() time.Duration
  DeliveryMethod(Callback)
	unique()
}
// STOPIFACE OMIT

// HIDDEN STRUCT IMPLEMENTING INTERFACE
//

// STARTSTRUCT OMIT
type order struct {
	taype string
	time  time.Duration
  cb Callback
}
// STOPSTRUCT OMIT

// METHODS
//

func (o *order) unique() {}

func (o *order) Type() string {
	return o.taype
}

func (o *order) Time() time.Duration {
	return o.time
}

func (o *order) DeliveryMethod(cb Callback) {
	o.cb = cb
}

// PACKAGE FUNCTIONS
//

// STARTNEW OMIT
func New(taype string, time time.Duration) Order {
	return &order{taype: taype, time: time}
}
// STOPNEW OMIT
