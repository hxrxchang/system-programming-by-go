package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	// go を付ける場合と付けない場合で time で実行すると違いがわかる
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(time.Second)
}
