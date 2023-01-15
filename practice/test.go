package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const s string = "constant"

func main() {
	fmt.Println("go" + "lang")
	fmt.Println("go \n" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println(true)
	fmt.Println(true && false)
	fmt.Println(true || false)

	/////VARIABLES+CONSTANTS

	fmt.Println(s + "\t" + "declare")
	const n = 50000
	fmt.Println(n)
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(d))
	fmt.Println("///////")

	//////LOOP

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println(i)
	fmt.Println("//////")
	for n := 0; n <= 10; n++ {
		if n == 5 {
			// break
		} else if n%2 == 1 {
			fmt.Println(n, "Odd")
			continue
		}
		fmt.Println(n, "Even")
	}
	fmt.Println("//////")

	//////SWITCH CASE
	mark := 89.9
	switch {
	case mark <= 100 && mark >= 90:
		fmt.Println("A+")
	case mark < 90 && mark >= 85:
		fmt.Println("A")
	case mark < 85 && mark >= 80:
		fmt.Println("B+")
	}

	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	fmt.Println(time.Now().Hour())

	findType := func(n interface{}) {
		switch t := n.(type) {
		case bool:
			fmt.Println("BOOL")
		case int:
			fmt.Println("INT")
		default:
			fmt.Printf("%T", t)
		}
	}
	findType("hello")
	fmt.Println("\n //////")

	////////Arrays

	var arr [5]string
	fmt.Println(len(arr))
	fmt.Println((arr))

	var a [5]int
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		if arr[i] == "" {
			fmt.Println("Empty")
			arr[i] = "A" + strconv.Itoa(i)
		}
	}
	fmt.Println((arr))

	var arr2d [3][2]int
	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			arr2d[i][j] = i + j
		}
	}

	fmt.Println("array2d:", arr2d)
	fmt.Println("row lenth :", len(arr2d))
	fmt.Println("row lenth :", len(arr2d[0]))

	//// SLICES

	s := make([]string, 5)
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		s[i] = "Slice" + strconv.Itoa(i)
	}
	fmt.Println("S", s)

	s = append(s, "Sapp1")
	fmt.Println(s, len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy C:", c)
	l := s[4:]
	fmt.Println(l)

	s2d := make([][]int, 3)
	fmt.Println(len(s2d))

	for i := 0; i < len(s2d); i++ {
		innerlen := i + 1
		s2d[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			s2d[i][j] = i + j
		}
	}
	fmt.Println("2d: ", s2d)

	/////MAPS

	m:=make(map[string]int)
	m["1"]=3
	m["0"]=2
	fmt.Println(m)

	// v1:=m[1]
	delete(m,"0")
	fmt.Println(m)

	fmt.Println(m["0"])
	_, present:=m["1"]
	fmt.Println(present)

	new:=map[string]string {"name": "ayesha","id":"C 1328"}
	fmt.Println(new)

	//// RANGE

	m["a"]=12
	m["ww"]=122
	for _, num:=range m{
		fmt.Println("m",num)
	}
	for i, num:=range m{
		fmt.Printf("%s : %d \n",i,num)
	}



}
