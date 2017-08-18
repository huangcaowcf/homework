package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com:80", "target host")
)

func handleConn(conn net.Conn) {
	//建立到目标服务器的连接
	//defer conn.Close()
	//var remote net.Conn
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println(*target)
	remote, err := net.Dial("tcp", *target)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	wg.Wait()
	//go 接收客户端的数据发送到remote
	//go 接收remote的数据，发送到客户端
	//等待连接关闭
}

func main() {
	//建立监听
	addr := "10.20.64.11:8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	//defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}
