// Concurrency1
package main

import (
	"fmt"
)

func main() {
	// testGoRoutine()
	// FirstChannel()
	BufferedChannel()
	
}

func testGoRoutine() {
	//Running in different memories
	LoopMe := func (wrd string) () {
		for i:=1; i<=5; i++ {
		fmt.Println(wrd)
		}
	}
	go LoopMe("Hi from here!") //Called as a Go Routine
	
	LoopMe("Hi from another process")
}

func BufferedChannel() {
	//Example on creating buffered Channel
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
