// main.go
package main

import (
	"fmt"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c //receive from c 
	fmt.Println(x, y, x+y)

	fmt.Println("For Buffered Channels")
	//Buffered Channels
	cd := make(chan int, 2)
	cd <- 1
	cd <- 2
	fmt.Println(<-cd)
	fmt.Println(<-cd)

	//Buffered channel 长度为1，只有channel中的值被读取，后续的值才能无阻塞的存进channel。
	cde := make(chan int, 1)
	cde <- 1
	fmt.Println(<-cde)
	cde <- 2
	fmt.Println(<-cde)

	fmt.Println("For Range and Close")
	//For use Range and Close of channel
	cx := make(chan int, 10)
	go fibonacci(cap(cx), cx)
	for i := range cx {
		fmt.Println(i)
	}

}
