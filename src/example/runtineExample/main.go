// main.go
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //cp把时间片让给别人，下次某个时候继续恢复执行goroutine.
		fmt.Println(s)
	}
}

func main() {
	go say("world") //open a new go runtines 
	say("hello")
}
