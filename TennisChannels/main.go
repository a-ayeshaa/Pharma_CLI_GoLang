package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init(){
	rand.Seed(time.Now().UnixNano())
}

func main(){
	court:=make(chan int) 

	wg.Add(2)

	go player("A",court)
	go player("B",court)

	court<-1
	wg.Wait()
}

func player(name string, court chan int){
	defer wg.Done()

	for {
		ball, ok:= <-court
		// fmt.Println(<-court)
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}

		n:=rand.Intn(100)
		// fmt.Println(n)

		if n%13==0{
			fmt.Println(n)
			fmt.Printf("Player %s Missed\n",name)

			close(court)
			return
		}

		fmt.Printf("Player %s Hit %d \n",name,ball)
		ball++

		court<-ball
	}
}