package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer1 bytes.Buffer
	buffer2 := bytes.NewBuffer([]byte{0x10, 0x20, 0x30})
	buffer3 := bytes.NewBufferString("初期文字列")
	fmt.Println(buffer1)
	fmt.Println(buffer2)
	fmt.Println(buffer3)
}
