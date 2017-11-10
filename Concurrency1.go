// Concurrency1
package main

import (
	"fmt"
	"time"
)

func main() {
	// testGoRoutine()
	// BasicChannel()
	// FirstChannel()
	// BufferedChannel()
	// channelSynchronization()
	// ChannelDirection()
	// goSelect()
	// SelectNonBlocking()
	// CloseChannels()
	RangeChannels()
}

func testGoRoutine() {
	//Running GoRoutine in seperate memories
	LoopMe := func (wrd string) () {
		for i:=1; i<=5; i++ {
		fmt.Println(wrd)
		}
	}
	go LoopMe("Hi from here!") //Called as a Go Routine
	
	LoopMe("Hi from another process")
}

func BasicChannel() {
	//Basic example of creating a channel, passing string to channel via Go Routine.
	chn := make(chan string)
	go func() { chn <- "ping" }()
	msg := <-chn
	fmt.Println(msg)
}

func FirstChannel() {
	//Basic channel routine
	sum := func(numbers []int, c chan int) {
		var total int
		for _, v:= range numbers {
			total += v
		}
		c <- total //send total to c
	}
	
	nos := []int{4,3,6,1,7,5,2,8}
	chn := make(chan int)
	go sum(nos[:len(nos)/2],chn)
	go sum(nos[len(nos)/2:],chn)
	go sum(nos, chn)
	
	x,y,z := <-chn, <-chn, <-chn
	fmt.Println(x, y, z, x+y)
}

func BufferedChannel() {
	//Example on creating buffered Channel. Values are transmitted sequentially & received too
	c := make(chan int, 3) //Second argument is the length of buffer
	c <- 6
	c <- 3
	c <- 5
	//c3 := func() { c <-5 }
	//go c3()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	
}

func channelSynchronization() {
	//Synchronize execution across go routines.
	
	worker := func(notify chan bool) {
				fmt.Print("Working....")	
				time.Sleep(time.Second)
				fmt.Println("Done")
				
				notify <- true
			}
		
	done := make(chan bool, 1)	//Channel used to notify that this function is completed. Sends a value to notify
	go worker(done) //Starts worker goroutine, passing a channel where it will be notified if its done
	<- done		//Worker wont execute until it is returned
}

func ChannelDirection() {
	//Direction can be set. Channel is meant to only send or receive values can be done if they are parameters
	pingfun := func(pings chan <- string, msg string) {
				pings <- msg
			}
	
	pongfun := func(pings <- chan string, pongs chan <- string) {
			msg := <- pings
			pongs <- msg
	}
	
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	
	pingfun(ping, "Testing channel direction")
	pongfun(ping, pong)
	fmt.Println(<-pong)
}

func goSelect() {
	//Select will wait & execute the particular case, whose channel has completed its operation
	chn1 := make(chan string)
	chn2 := make(chan string)
	
	//1st Go routine for Chn1
	go func() {
		time.Sleep(time.Second * 3)
		chn1 <- "Channel 1"
	}()
	
	//2nd Go routine for Chn2
	go func() {
		time.Sleep(time.Second * 1)
		chn1 <- "Channel 2"
	}()
	
	for i:=0; i<2; i++ {
		select {
			case msg1 := <-chn1:
				fmt.Println("Received from ", msg1)
			case msg2 := <-chn2:
				fmt.Println("Received from ", msg2)
		}
	}
}

func SelectNonBlocking() {
//Demonstrating Non Blocking send/receive. If channel does not receive or send, default block is executed
	messages := make(chan string)
	//signals := make(chan string)
	
	go func() {	messages <- "this is cool" } ()
	//time.Sleep(time.Second * 1)
	for {	//infinite loop
		select {
			case msg := <- messages:
				fmt.Println("Received from Messages", msg)
				return
			default:
				fmt.Println("No message received")
		}
	}
}

func CloseChannels() {
	//This shows how to Close a particular channel, once the work is done
	jobs := make(chan int, 5)
	done := make(chan bool)
	
	go func() {		//In this Go Routine, we are checking if the jobs channel is open & receiving data
		for {
			data, IsOpen := <- jobs
			if IsOpen {
				fmt.Println("Data received", data)
			} else {	//Channel is closed
				fmt.Println("Data receiving complete!")
				done <- true
				return	
			}
		}
	}()
	
	//Transmitting data over jobs channel
	for i:=0; i<3; i++ {
		jobs <- i
		fmt.Println("Data Sent", i)
		time.Sleep(time.Second * 1)
	}
	close(jobs)		//Close the channel, once all data is sent
	fmt.Println("Channel Closed!")
}

func RangeChannels() {
	//Passing multiple values to a channel & iterating through it
	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	queue <- "three"
	close(queue)	
	
	for elem:=range queue {
		fmt.Println(elem)
	}
}