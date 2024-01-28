package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("Channels in golang")

	// channels are used for communication between two go routines

	myCh := make(chan int, 2) // 2 is buffer
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	// consider channel as box, <-chan means value going out of channel, chan<- means value getting inside the channel
	// receive only
	go func(ch <-chan int, wg *sync.WaitGroup){
		val, isClosed := <-myCh
		fmt.Println(isClosed) // check whether the 0 is val or default closed channel value
		fmt.Println(val)
		wg.Done()
	}(myCh,wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup){
		myCh<-5
		myCh<-6 // will one more listener for this, or else define a buffer
		close(myCh) // will give 0 as output when listened
		wg.Done()
	}(myCh,wg)

	wg.Wait()
}