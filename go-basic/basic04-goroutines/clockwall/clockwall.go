package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	for _, v := range os.Args[1:] {
		keyValue := strings.Split(v, "=")
		go connTCP(keyValue[1])
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func connTCP(uri string) {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
