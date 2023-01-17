package main

import (
	"fmt"
	"time"
)

func main(){
	//GOROUTINES
	f("synchronous")


	go f("123")
	go f("asynchronous")

	func (msg string)  {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	
	//CHANNELS
	messages:=make(chan string)

	// go f(<-messages)
	go func (){
		messages<-"ping"
	}()
	
	msg:=<-messages
	fmt.Println(msg)

	////CHANNEL BUFFERING

	message:=make(chan string,2)
	message<-"buffered"
	message<-"channel"
	fmt.Println(<-message)
	fmt.Println(<-message)

	done:=make(chan bool,1)
	go worker(done)
	<-done

	//CHANNEL DIRECTIONS

	pings:=make(chan string,1)
	pongs:=make(chan string,1)
	ping(pings,"passed message")
	pong(pings,pongs)
	fmt.Println(<-pongs)

	//SELECT

	c1:=make(chan string)
	c2:=make(chan string)

	go func ()  {
		time.Sleep(1*time.Second)
		c1<-"one"
	}()

	go func ()  {
		time.Sleep(1*time.Second)
		c2<-"two"
	}()

	for i:=0;i<2;i++{
		select{
		case msg1:=<-c1:
			fmt.Println("received",msg1,i)
		case msg2:=<-c2:
			fmt.Println("received",msg2,i)
		}
	}
}