package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		done <- true
	}()
	<- done
	fmt.Println("all tasks are finished")
}
