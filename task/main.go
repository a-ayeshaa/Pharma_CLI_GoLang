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

	fmt.Println(m.get(2))

	newmed := Medicine{10, "Napa10", 1000}
	fmt.Println(m.add(newmed))
	fmt.Println(m.getAll())
	m.delete(0)
	fmt.Println(m.getAll())

	medupdate:=Medicine{10,"Napa10",1500}
	fmt.Println(m.update(medupdate))
	fmt.Println(m.getAll())

}
