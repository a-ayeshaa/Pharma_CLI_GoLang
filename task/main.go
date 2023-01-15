package main

import (
	"fmt"
	"strings"
)

func main() {

	m := Medicine{}
	meds := m.getAll()

	fmt.Println("The available medicines are : ")
	fmt.Printf("%s \n", strings.Repeat("-", 42))
	fmt.Printf("| %10s | %10s | %12s  |\n", "Id", "Name", "Price")
	fmt.Printf("%s \n", strings.Repeat("-", 42))
	for _, med := range meds {
		fmt.Printf("| %10d | %10s | %10d tk |\n", med.id, med.name, med.price)
	}
	fmt.Printf("%s \n", strings.Repeat("-", 42))

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
			fmt.Println(m.getAll())
		case 2:
			var index int
			fmt.Println("Enter the ID of the medicine you want to delete: ")
			fmt.Scanln(&index)
			m.delete(index)
			fmt.Println(m.getAll())
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
			fmt.Println(m.getAll())
		case 4:
			var index int
			fmt.Println("Enter the Id of the medicine :")
			fmt.Scanln(&index)
			fmt.Println(m.get(index))
		case 5:
			end = true
		}

	}

}
