package main

import (
	"fmt"
	"os"
)
func mayPanic(){
	panic("a problem")
}
func main(){
	f:=createFile("/home/pathao/Documents/GoLangIntro/defer.txt")
	defer closeFile(f)
	writeFile(f)

	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered error :",r)
		}
	}()
	mayPanic()
	_,err:=os.Create("/home/pathao/Documents/GoLangIntro/defer1.txt")
	if err!=nil{
		panic(err)
	}
}

func createFile(p string) *os.File{
	fmt.Println("creating")
	f,err:=os.Create(p)
	if err!=nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File){
	fmt.Println("writing")
	fmt.Fprintln(f,"data")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	err:=f.Close()
	if err!=nil{
		fmt.Fprintf(os.Stderr,"error: %v\n",err)
		os.Exit(1)
	}
}