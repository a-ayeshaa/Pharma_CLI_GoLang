package main

import (
	"fmt"
	"time"
)

func f(from string){
	for i:=0;i<3;i++{
		fmt.Println(from,":",i)
	}
}
//CHANNEL BUFFERING
func worker (done chan bool) {
	// <-done
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done<-true
}

//CHANNEL DIRECTIONS
func ping(pings chan<-string, msg string){
	pings<-msg
}

func pong(pings<-chan string, pongs chan<- string){
	msg:=<-pings
	pongs<-msg
}