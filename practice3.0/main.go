package main

import (
	"fmt"
	"time"
)

func main() {
	//GOROUTINES
	f("synchronous")

	go f("123")
	go f("asynchronous")

	func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)

	//CHANNELS
	messages := make(chan string)

	// go f(<-messages)
	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)

	////CHANNEL BUFFERING

	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"
	fmt.Println(<-message)
	fmt.Println(<-message)

	done := make(chan bool, 1)
	go worker(done)
	<-done

	//CHANNEL DIRECTIONS

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	//SELECT

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1, i)
		case msg2 := <-c2:
			fmt.Println("received", msg2, i)
		}
	}

	//Timeouts

	cc1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		cc1 <- "result 1"
	}()

	select {
	case res := <-cc1:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 1")
	}

	//NON BLOCKING CHANNEL OPERATIONS

	msgs := make(chan string)
	// signals:=make(chan bool)
	go func() {
		msgs <- "hello"
	}()
	time.Sleep(1 * time.Second)

	select {
	case ms:=<-msgs:
		fmt.Println("received msg", ms)
		// fmt.Println("received msg", <-msgs)
	default:
		fmt.Println("no message received")
	}

	time.Sleep(1 * time.Second)

	//CLOSING CHANNELS

	jobs := make(chan int, 5)
	done1 := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done1 <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	fmt.Println("sent all jobs")

	<-done1
}
