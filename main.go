package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	//fmt.Printf("%d", <-ch)
	go tel(ch, quit)

	ok := true
	//i := 0

	for ok {
		//if i, ok = <-ch; ok {
		//	fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		//}
		select {
		case i := <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
	//time.Sleep(2e9)
}

func tel(ch chan int, quit chan bool) {
	for i := 0; i < 20; i++ {
		ch <- i
	}
	//close(ch)

	quit <- true
}

func f1(in chan int) {
	fmt.Println(<-in)
}
