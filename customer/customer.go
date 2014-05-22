package customer

import (
	"github.com/vaskoz/coffee-shop/order"
  "fmt"
)

type Customer interface {
	PlaceOrder() order.Order
  EnjoyBeverage(order.Order, interface{}) string
	Id() int
}

type customer struct {
	id int
}

func (c *customer) PlaceOrder() order.Order {
	return order.New("latte", 2)
}

func (c *customer) Id() int {
	return c.id
}

func (c *customer) EnjoyBeverage(order order.Order, provider interface{}) string {
  return fmt.Sprintf("Customer %d says Yum and thanks to %v", c.id, provider)
}

func (c *customer) String() string {
  return fmt.Sprintf("Customer %d", c.id)
}

func New(id int) Customer {
	return &customer{id: id}
}

func RandomGroupOf(n int) []Customer {
	customers := make([]Customer, n)
	for i := 0; i < n; i++ {
		customers[i] = New(i+1)
	}
	return customers
}
