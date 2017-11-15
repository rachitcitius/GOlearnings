// GoRoutines_Cntd
package main

import (
	"fmt"
	"time"
)

func main() {
	// TimerStop()
	// TickerStart()
	WorkerPool()
}

func TimerStop() {
	
	//Timer provides a channel ".C" to notify the wait. Timer can be stopped before they are expired
	
	timer1 := time.NewTimer(time.Second * 2)
	<- timer1.C		//blocks until the timer expired
	fmt.Println("Timer 1 Expired")
	
	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 Expired")
	}()
	
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 Stopped")
		}
}

func TickerStart() {
	//Ticker to do something at regular intervals
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticked at ", t)
		}
	} ()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker Stopped!")
}

func WorkerPool() {
	//creating worker pool 
	worker := func(id int, jobs<-chan int, results chan<-int) {
		for j:=range jobs {
			fmt.Println("Worker", id, "started job", j)
			time.Sleep(time.Second)
			fmt.Println("Worker", id, "finished job", j)
			results <- j*2
		}
	}
	
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	for w:=1;w<=3;w++ {
		//getting workers to start working
		go worker(w, jobs, results)
	}
	
	for i:=1;i<=5;i++ {
		//Sending job
		jobs <- i * 10
	}
	close(jobs)
	
	for a:=1;a<=5;a++ {
		//collecting results
		<-results
	}
}