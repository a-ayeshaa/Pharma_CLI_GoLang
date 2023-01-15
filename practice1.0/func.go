package main

import (
	"fmt"
	
	"unicode/utf8"
)

func add(a int,b int) []int{
	arr:=make([]int,2)
	arr=append(arr,a)
	arr=append(arr,b)
	return arr
}

func multipleval() (int,int,int){
	return 2,3,4
}

func sum(nums ...int){
	fmt.Println(nums)
	fmt.Println(len(nums))
}

func nextSeq() func() int {
	i:=0
	return func() int{
		i++
		return i
	}
}

func fact(n int) int{
	if n==0{
		return 1
	} else {
		return n*fact(n-1)
	}
}

func zeroval(ival int) {
	ival=0
}

func zeroptr(ival *int){
	*ival=0
}
func main(){
	// fmt.Println("1+1",add(1,1))
	_,_,b:=multipleval()
	fmt.Println(b)
	sum(1,2,3,4)
	sum(1)
	fmt.Println(add(1,2))
	next:=nextSeq()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	i:=10
	fmt.Println("initial:",i)
	zeroval(i)
	fmt.Println("after zeroval:",i)
	zeroptr(&i)
	fmt.Println("after zeroptr:",i)

	///// STRINGS AND RUNES

	// var s string= "ayesha akhtar"
	const s= "$%^????}}}"
	// var s string="สวัสดี"

	fmt.Println("Len :",len([]rune(s)))

	fmt.Println("Rune count: ",utf8.RuneCountInString(s))

	for index, char:=range s{
		fmt.Printf("%#U  , %d \n\n\n",char,index)
	}

	//two variables used in for loop
	for i,w:=0, 0; i<len(s); i+=w{
		fmt.Println(i,w,s[i:])

		runevalue,width:=utf8.DecodeLastRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n",runevalue,i)
		w=width

	}

}