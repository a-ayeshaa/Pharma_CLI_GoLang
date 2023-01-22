package main

import (
	"io"
	"log"
	// "sync"
	"sync/atomic"
)

const (
	maxGoroutines =25
	pooledResources=2
)

type dbConnection struct{
	ID int32
}

func (dbConn *dbConnection) Close() error{
	log.Println("CLose: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer,error){
	id:= atomic.AddInt32(&idCounter,1)
	log.Println("Create:New Connection",id)
	return &dbConnection{id},nil
}

func main(){
	// var wg sync.WaitGroup
	// wg.Add(maxGoroutines)

	// pool:= Pool{}
	// p,err:=pool.New(createConnection,pooledResources)
	// if err!=nil{
	// 	log.Println(err)
	// }

	// for query:=0;query<maxGoroutines;query++{
	// 	go func(q int){
	// 		performQueries(q,p)
	// 		wg.Done()
	// 	}(query)
	// }
}