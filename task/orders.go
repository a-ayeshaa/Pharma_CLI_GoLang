package main

var order []Order

type Order struct {
	id         int
	username   string
	totalprice int
}

type IOrder interface {
	addOrder(o Order) Order
}

func (or Order) addOrder(o Order) Order {
	var id int
	if len(order) == 0 {
		id = 0
	} else {
		id = order[len(order)-1].id + 1
	}
	o.id = id
	order = append(order, o)
	cart=make([]Cart, 0)
	return order[id]

}
