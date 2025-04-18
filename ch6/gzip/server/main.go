package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}

func isGzipAcceptable(req *http.Request) bool {
	acceptEncodingHeaders := strings.Join(req.Header["Accept-Encoding"], ",")
	i := strings.Index(acceptEncodingHeaders, "gzip")
	return i != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		req, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			neterr, ok := err.(net.Error)
			if ok && neterr.Timeout() {
				fmt.Println("timeout")
				break
			} else if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		res := http.Response{
			StatusCode: 200,
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header: make(http.Header),
		}
		if isGzipAcceptable(req) {
			content := "hello world(gzipped)\n"
			var buffer bytes.Buffer
			writer := gzip.NewWriter(&buffer)
			io.WriteString(writer, content)
			writer.Close()
			res.Body = io.NopCloser(&buffer)
			res.ContentLength = int64(buffer.Len())
			res.Header.Set("Content-Encoding", "gzip")
		} else {
			content := "hello world\n"
			res.Body = io.NopCloser(strings.NewReader(content))
			res.ContentLength = int64(len(content))
		}
		res.Write(conn)
	}
}
