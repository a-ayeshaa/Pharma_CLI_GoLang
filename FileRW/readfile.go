package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error){
	if e!=nil{
		panic(e)
	}
}

func main(){
	dat,err:=os.ReadFile("/home/pathao/Documents/GoLangIntro/defer.txt")
	check(err)
	fmt.Println(string(dat))

	f,err:=os.Open("/home/pathao/Documents/GoLangIntro/defer.txt")
	check(err)

	b1:=make([]byte,2)
	n1,err:=f.Read(b1) //returns the number of bytes read
	check(err)
	fmt.Printf("%d bytes: %s\n",n1,string(b1[:n1]))

	o2,err:=f.Seek(2,0)
	check(err)
	b2:=make([]byte,25)
	n2,err:=f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d:",n2,o2)
	fmt.Printf("%v\n",string(b2[:n2]))

	o3,err:=f.Seek(2,0)
	check(err)
	b3:=make([]byte,2)
	n3,err:=io.ReadAtLeast(f,b3,2)
	fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))

	r4:=bufio.NewReader(f)
	b4,err:=r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes : %s\n",string(b4))
	f.Close()

	
	//WRITING TO FILES

    d1 := []byte("hello\ngo\n")
    err1:= os.WriteFile("/home/pathao/Documents/GoLangIntro/dat1", d1, 0644)
    check(err1)

    f1, err := os.Create("/home/pathao/Documents/GoLangIntro/dat2")
    check(err)

    defer f1.Close()

    d2 := []byte{115, 111, 109, 111, 10}
    n22, err := f1.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n22)

    n32, err := f1.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n32)

    f1.Sync()

    w := bufio.NewWriter(f1)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()


}