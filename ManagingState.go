// ManagingState
//Accessing counter state from across multiple goroutines using various techniques
package main

import (
	"fmt"
	"time"
	"sync/atomic"
	"math/rand"
	"sync"
)

func main() {
	// UsingAtomicCounter()
	UsingMutex()
}

func UsingAtomicCounter() {
	//Using Atomic Counters instead of channel for concurrency 
	var ops uint64 = 0
	
	//Start 50 goroutines that each increment the counter once a ms
	for i:=0; i<50; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)	//Incrementing the counter by 1
			
			time.Sleep(time.Millisecond)
		}()
	}
	
	time.Sleep(time.Second)
	//Reading the incremented values
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

func UsingMutex() {
	//Using Mutex we will access state across other goroutines
	var state =  make(map[int]int)	//State will be a map
	
	var mutex = &sync.Mutex{}	//This mutex will sync access to state
	
	//Track Read & Write operations
	var readOps uint64 = 0
	var writeOps uint64 = 0
	
	//Starting 100 goroutines for repeated reads of the state, once per ms in each goroutine
	for r:=0; r<100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()	//Locking to ensure exclusive access to the state
				total += state[key]	//Read the state
				mutex.Unlock()	//Unlock
				atomic.AddUint64(&readOps, 1)	//Increament the read count
				
				time.Sleep(time.Millisecond)
			}
		}()
	}
	
	//10 goroutines to write, similar to reads
	for w:=0; w<10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value	//Writing to the state
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				
				time.Sleep(time.Millisecond)
			}
		}()
	}
	
	time.Sleep(time.Second)
	
	fmt.Println("Read Ops:", readOps)
	fmt.Println("Write Ops:", writeOps)
	
	mutex.Lock()
	fmt.Println("State:", state)
	mutex.Unlock()
}