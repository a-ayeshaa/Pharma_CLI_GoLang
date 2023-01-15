package main

import (
	"fmt"
	"time"
	// P "functions"
)

func main() {
	///person
	fmt.Println((person{"Ayesha", 23}))

	fmt.Println(&person{name: "Nico", age: 3})

	fmt.Println(NewPerson("Pixie", 2))
	s := newPerson("Cookie", 2)
	newp := s
	fmt.Println(newp, s)

	newp.age = 4

	fmt.Println(newp, s)

	///rectangle
	r := rectange{width: 5, height: 6}

	a := allrectange(r)
	fmt.Println("Area:", a.area())
	fmt.Println("Perimeter:", a.perimeter())

	//////Struct embedding

	co := base{}
	co.num = 22
	fmt.Println(co)

	fmt.Println(co.describe())

	embedded := container{}
	embedded.base = co
	// embedded.num=232
	embedded.str = "Embedded struct has been implemented"
	fmt.Printf("embedded values are : co {num: %v, str: %v} \n", embedded.num, embedded.str)

	////GENERICS

	var m = map[int]string{1: "2", 2: "4", 4: "8", 6: "16"}
	fmt.Println("keys:", Mapkeys(m))
	fmt.Println("keys:", Mapvals(m))
	fmt.Println(len(m))

	// _ = Mapkeys[int,string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(20)
	fmt.Println("list:", lst.GetAll())

	/////ERRORS

	for _, i := range []int{4, 40} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)

		} else {
			fmt.Println("f1 worked", r)
		}
	}

	for _, i := range []int{4, 40} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed", e)

		} else {
			fmt.Println("f1 worked", r)
		}
	}

	_, e := f2(40)
	if ae, ok := e.(*argError); ok {
		// fmt.Println(ok)
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	f("direct")

	go f("goroutine")

	go func (msg string)  {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println(time.Second)
	fmt.Println("done")
}
