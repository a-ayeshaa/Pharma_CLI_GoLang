package main

import (
	"fmt"
	"strings"
)

type Cart struct {
	id         int
	name       string
	totalprice int
	quantity   int
}

var cart []Cart

func cartlist() []Cart {
	return cart
}

type ICart interface {
	addtoCart(c Cart) Cart
	remove(id int) bool
	printcart(carts []Cart)
	confirm() bool
}

func (ca Cart) addtoCart(c Cart) Cart {
	cart = append(cart, c)
	return c
}

func (ca Cart) remove(id int) bool {
	for i, cartval := range cart {
		if cartval.id == id {
			cart = append(cart[:i], cart[i+1:]...)
			return true
		}

	}
	return false
}

func (ca Cart) printcart(carts []Cart) {
	if len(carts) == 0 {
		fmt.Println("Cart is empty")
	} else {
		var total int = 0
		fmt.Println("Your shopping cart : ")
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		fmt.Printf("| %10s | %10s | %12s  |\n", "Name", "Quantity", "Price")
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		for _, cart := range cart {
			total += cart.totalprice
			fmt.Printf("| %10s | %10d | %10d tk |\n", cart.name, cart.quantity, cart.totalprice)
		}
		fmt.Printf("%s \n", strings.Repeat("-", 42))
		fmt.Printf("Total Price : %d tk \n", total)
		fmt.Printf("%s \n", strings.Repeat("-", 42))
	}
}

func (ca Cart) confirm() bool{
	if len(cart)!=0{
		var total int=0
		for _,val:=range cart{
			total+=val.totalprice
		}

		o:=Order{}
		order:=Order{
			username: "customer",
			totalprice: total,
		}
		fmt.Println(o.addOrder(order))
		return true
	}
	return false
}

func searchMed(id int) Cart {
	m := Medicine{}
	val := m.get(id)
	newcart := Cart{
		id:         val.id,
		name:       val.name,
		totalprice: val.price,
		quantity:   0,
	}
	return newcart
}
