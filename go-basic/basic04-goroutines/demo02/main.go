package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var title = "default location"

func main() {
	var port = flag.String("port", "", "请输入端口")

	if *port == "8000" {
		title = "ShangHai"
	} else {
		title = "NewYork"
	}
	fmt.Printf("port: %s", *port)
	// 创建了一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 获取端口的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format(title+" time 15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
