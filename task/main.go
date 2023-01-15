package main

import (
	"fmt"
	"strings"

)

func printlist(meds []Medicine) {
	fmt.Println("The available medicines are : ")
	fmt.Printf("%s \n", strings.Repeat("-", 42))
	fmt.Printf("| %10s | %10s | %12s  |\n", "Id", "Name", "Price")
	fmt.Printf("%s \n", strings.Repeat("-", 42))
	for _, med := range meds {
		fmt.Printf("| %10d | %10s | %10d tk |\n", med.id, med.name, med.price)
	}
	fmt.Printf("%s \n", strings.Repeat("-", 42))
}

func main() {

	m := Medicine{}
	c := Cart{}
	meds := m.getAll()
	printlist(meds)

	var user string
	fmt.Println("Enter Username:")
	fmt.Scanln(&user)
	if strings.ToLower(user) == "admin" {
		var end bool = false
		for !end {
			var choice int
			fmt.Println("Admin, you have the following options to execute: \n 1. Add Medicine \n 2. Delete Medicine \n 3. Edit Medicine \n 4. Get Medicine \n 5. Exit")
			fmt.Println("Choose an option: ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				var index int = medlist[len(medlist)-1].id + 1
				var price int
				var name string
				fmt.Println("Set Name: ")
				fmt.Scanln(&name)
				fmt.Println("Set Price: ")
				fmt.Scanln(&price)

				newmed := Medicine{index, name, price}
				fmt.Println(m.add(newmed))
				printlist(m.getAll())

			case 2:
				var index int
				fmt.Println("Enter the ID of the medicine you want to delete: ")
				fmt.Scanln(&index)
				m.delete(index)
				printlist(m.getAll())

			case 3:
				var index int
				var price int
				var name string
				fmt.Println("Enter the Id of the medicine :")
				fmt.Scanln(&index)
				fmt.Println("Enter the Name of the medicine :")
				fmt.Scanln(&name)
				fmt.Println("Enter the Price of the medicine :")
				fmt.Scanln(&price)

				medupdate := Medicine{index, name, price}
				fmt.Println(m.update(medupdate))
				printlist(m.getAll())

			case 4:
				var index int
				fmt.Println("Enter the Id of the medicine :")
				fmt.Scanln(&index)
				fmt.Println(m.get(index))
			case 5:
				end = true
			}

		}
	} else if strings.ToLower(user) == "customer" {
		var end bool = false
		for !end {
			var choice int
			fmt.Println("Customer, you have the following options to execute: \n 1. Add Medicine to cart \n 2. Remove from Cart \n 3. Check Cart \n 4. Confirm Order \n 5. Exit")
			fmt.Println("Choose an option: ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				var id int
				var q int
				fmt.Println("Enter ID :")
				fmt.Scanln(&id)
				add:=searchMed(id)
				fmt.Println("Enter Quantity :")
				fmt.Scanln(&q)
				add.totalprice=add.totalprice*q
				add.quantity=q
				// cart:=Cart{1,"Napa",200,2}
				c.addtoCart(add)
				fmt.Println(cartlist())
			case 2:
				var id int
				fmt.Println("Enter ID:")
				fmt.Scanln(&id)
				c.remove(id)
				fmt.Println(cartlist())
			case 3:
				c.printcart(cart)
			case 4:
				val:=c.confirm()
				if val {
					fmt.Println("Your order has been confirmed, order again!")
				} else {
					fmt.Println("Your cart is empty.")
				}
			case 5:
				end=true
			}
		}
	}

}
