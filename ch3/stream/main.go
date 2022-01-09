package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	header := bytes.NewBufferString("----- HEADER -----\n")
	content := bytes.NewBufferString("Example of io.MultiReader\n")
	footer := bytes.NewBufferString("----- Footer -----\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
}
