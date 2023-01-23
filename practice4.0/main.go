package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs<-chan int,results chan<-int){
	for j:=range jobs{
		fmt.Println("worker",id,"started job",j)
		time.Sleep(time.Second)
		fmt.Println("worker",id,"finished job",j)
		results<-j*2
	}
}

func work(id int){
	fmt.Printf("WOrker %d starting\n",id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n",id)
}

func main(){
	queue:=make(chan string,2)
	queue<-"one"
	queue<-"two"
	close(queue)
	// fmt.Println(<-queue)
	// fmt.Println(<-queue)

	for elem:=range queue{
		fmt.Println("QUEUE:",elem)
	}

	///TIMERS//

	///Worker pools

	const numjobs=5
	jobs:=make(chan int,numjobs)
	results:=make(chan int, numjobs)

	for w:=1;w<=3;w++{
		go worker(w,jobs,results)
	}
	for j:=1;j<=numjobs;j++{
		jobs<-j
	}
	close(jobs)
	for a:=1;a<=numjobs;a++{
		<-results
	}

	//WAIT GROUPS
	fmt.Print("////////////\n")
	var wg sync.WaitGroup
	for i:=1;i<=5;i++{
		wg.Add(1)
		i:=i
		go func ()  {
			defer wg.Done()
			work(i)
		}()
	}

	wg.Wait()



}
