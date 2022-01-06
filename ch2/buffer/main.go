package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	var buffer bytes.Buffer
	var builder strings.Builder
	buffer.Write([]byte("bytes.Buffer example\n"))
	io.WriteString(&buffer, "bytes.Buffer example\n")
	builder.Write([]byte("strings.Buffer example\n"))
	fmt.Println(buffer.String())
	fmt.Println(builder.String())
}
