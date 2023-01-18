package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const(
	numberGoroutines =4
	taskLoad = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks:= make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr:=1;gr<=numberGoroutines;gr++{
		go worker(tasks,gr)
	}

	for post:=1;post<=taskLoad;post++{
		tasks<-fmt.Sprintf("Task : %d",post)
	}

	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int){
	defer wg.Done()

	for{
		tasks,ok:=<-tasks
		if !ok{
			fmt.Printf("Worker: %d : Shutting down \n",worker)
			return
		}

		fmt.Printf("Worker: %d :Stared %s \n",worker,tasks)

		sleep:=rand.Int63n(100)
		time.Sleep(time.Duration(sleep)*time.Millisecond)

		fmt.Printf("Worker: %d :Completed %s \n",worker,tasks)
	}
}