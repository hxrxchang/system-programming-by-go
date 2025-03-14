package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func main() {
	sendMessages := []string{
		"ASCII",
		"PROGRAMMING",
		"PLUS",
	}
	current := 0
	var conn net.Conn = nil

	for {
		var err error
		if conn == nil {
			conn, err = net.Dial("tcp", "localhost:8888")
			if err != nil {
				panic(err)
			}
			fmt.Printf("Access: %v\n", current)
		}

		req, err := http.NewRequest("POST", "http://localhost:8888", strings.NewReader(sendMessages[current]))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Accept-Encoding", "gzip")
		err = req.Write(conn)
		if err != nil {
			panic(err)
		}

		res, err := http.ReadResponse(bufio.NewReader(conn), req)
		if err != nil {
			fmt.Println("Retry")
			conn = nil
			continue
		}

		dump, err := httputil.DumpResponse(res, false)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		defer res.Body.Close()

		if res.Header.Get("Content-Encoding") == "gzip" {
			reader, err := gzip.NewReader(res.Body)
			if err != nil {
				panic(err)
			}
			io.Copy(os.Stdout, reader)
			reader.Close()
		} else {
			io.Copy(os.Stdout, res.Body)
		}

		current++
		if current == len(sendMessages) {
			break
		}
	}

	conn.Close()
}
