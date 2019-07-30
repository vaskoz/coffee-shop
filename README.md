[![Build Status](https://travis-ci.org/vaskoz/coffee-shop.svg?branch=master)](https://travis-ci.org/vaskoz/coffee-shop)
[![Coverage Status](https://coveralls.io/repos/github/vaskoz/coffee-shop/badge.svg?branch=master)](https://coveralls.io/github/vaskoz/coffee-shop?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/vaskoz/coffee-shop)](https://goreportcard.com/report/github.com/vaskoz/coffee-shop)
[![GoDoc](https://godoc.org/github.com/vaskoz/coffee-shop?status.svg)](https://godoc.org/github.com/vaskoz/coffee-shop)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.txt)

Use Cases:
===========

1. Customers enter, place orders, and only have so much patience for a
   busy line. Each customer gets in line, but they will leave if they
need to wait too long for their order.

2. Barista's fulfill orders to the best of their abilities. Their
   abilities range from 1.0 (which is perfect ability) to higher numbers
which are multiples of the optimal drink preparation time.

3. The store manager ensure that inventory is properly stocked and
   accounted for.

Running the program:
====================

The preferred way to run the program is after running `go build` in the
package.
```
./coffee-shop
main.go:30: please specify number, in seconds, for envvar COFFEE_SHOP_CLOSE_TIME
main.go:35: please specify number, in seconds, for envvar COFFEE_SHOP_SHUTDOWN
main.go:40: please specify number of customers for envvar COFFEE_SHOP_CUSTOMERS
main.go:45: please specify number of baristas for envvar COFFEE_SHOP_BARISTAS
```

The program requires 4 environment variables to be set before simulating
a coffee shop. Below is an example run of the program.

```
COFFEE_SHOP_CLOSE_TIME=60 COFFEE_SHOP_SHUTDOWN=5 COFFEE_SHOP_CUSTOMERS=50 COFFEE_SHOP_BARISTAS=2 ./coffee-shop
store.go:65: Customer 1 says Yum and thanks to Barista 1
store.go:65: Customer 2 says Yum and thanks to Barista 2
store.go:65: Customer 3 says Yum and thanks to Barista 1
^Cmain.go:63: I received a signal to close the store
store.go:65: Customer 5 says Yum and thanks to Barista 1
store.go:54: Store is closing
store.go:65: Customer 4 says Yum and thanks to Barista 2
store.go:65: Customer 6 says Yum and thanks to Barista 1
main.go:70: Store closed

```

`COFFEE_SHOP_CLOSE_TIME=60` means that the coffee shop is serving
customers for 60 seconds.

`COFFEE_SHOP_SHUTDOWN=5` means that the coffee shop tries to complete
serving existing customers for 5 seconds after the store closes.

`COFFEE_SHOP_CUSTOMERS=50` indicates there are 50 customers trying to be
coffee in the shop.

`COFFEE_SHOP_BARISTAS=2` indicates the number of baristas working on
making drinks. Each barista can only make one drink at a time.
