package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("ch3/file/main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
