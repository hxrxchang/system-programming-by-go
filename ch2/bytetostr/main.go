package main

import "fmt"

func main() {
	byteArray := []byte("ASCII")
	fmt.Println(byteArray)

	str := string([]byte{65, 83, 67, 73, 73})
	fmt.Println(str)
}
