package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("ch3/file/copy/main.go")
	if err != nil {
		panic(err)
	}
	newfile, err := os.Create("ch3/file/copy/copy.txt")
	if err != nil {
		panic(err)
	}
	io.Copy(newfile, file)
	file.Close()
	newfile.Close()
}
