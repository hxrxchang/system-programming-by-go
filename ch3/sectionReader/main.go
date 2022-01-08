package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	SectionReader := io.NewSectionReader(reader, 14, 2)
	io.Copy(os.Stdout, SectionReader)
}
