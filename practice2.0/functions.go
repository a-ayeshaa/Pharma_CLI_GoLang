package main

import (
	"errors"
	"fmt"
)

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {
	p := person{}
	p.name = name
	p.age = age

	return &p
}

func newPerson(name string, age int) person {
	p := person{}
	p.name = name
	p.age = age

	return p
}

/////

type rectange struct {
	width, height int
}

func (r rectange) area() int {
	return r.width * r.height
}

func (r rectange) perimeter() int {
	return r.width*2 + r.height*2
}

type allrectange interface {
	area() int
	perimeter() int
}

func calculate(a allrectange) (int, int) {
	return a.area(), a.perimeter()
}

///struct embedding

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

type container struct {
	base
	str string
}

////GENERICS

func Mapkeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Mapvals[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for _, num := range m {
		r = append(r, num)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

////ERRORS

func f1(arg int) (int, error) {
	if arg == 40 {
		return -1, errors.New("Cannot work with 40")
	} else {
		return arg + 3, nil
	}
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg==40 {
		return -1, &argError{arg, "cannot work with 40"}
	}
	return arg+3,nil
}

///////GOROUTINES

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}