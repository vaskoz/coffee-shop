package customer

import (
	"github.com/vaskoz/coffee-shop/item"
)

type Customer interface {
	PlaceOrder() item.Item
	RetrieveOrder(item.Item)
	Id() int
}

type customer struct {
	id int
}

func (c *customer) PlaceOrder() item.Item {
	return item.New("latte", 2)
}

func (c *customer) RetrieveOrder(item item.Item) {
	print(c.Id())
	println("Thanks for the delicious " + item.Type())
}

func (c *customer) Id() int {
	return c.id
}

func New(id int) Customer {
	return &customer{id: id}
}

func RandomGroupOf(n int) []Customer {
	customers := make([]Customer, n)
	for i := 0; i < n; i++ {
		customers[i] = New(i)
	}
	return customers
}
