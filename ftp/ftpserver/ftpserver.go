package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func handleConn(conn net.Conn) {
	//从conn里面读取一行内容
	//按空格分割指令和文件名
	//打开文件
	//读取内容
	//发送内容
	//关闭连接和文件
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fmt.Println(line)
	fields := strings.Fields(line)
	var cmd string
	var name string

	switch len(fields) {
	case 0:
		return
	case 1:
		cmd = fields[0]
	case 2:
		cmd = fields[0]
		name = fields[1]
	default:
		fmt.Println("Input error,please reinput:")
		conn.Write([]byte("Bad input,Input error,please reinput:\n"))
		return
	}

	switch {
	case cmd == "ls" && len(fields) == 1:
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Ls本当前目录:", dir)
		info, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range info {
			conn.Write([]byte(file.Name() + "\n"))
		}

	case cmd == "ls" && len(fields) == 2:
		info, err := ioutil.ReadDir(name)
		if err != nil {
			log.Println(err)
			conn.Write([]byte("Bad input,Input error,please reinput:\n"))
			return
		}
		fmt.Println("Ls目录:", name)
		for _, file := range info {
			conn.Write([]byte(file.Name() + "\n"))
		}

	case cmd == "pwd" && len(fields) == 1:
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		conn.Write([]byte(dir + "\n"))

	case cmd == "get" && len(fields) == 1:
		fmt.Println("Input error,please reinput:")
		conn.Write([]byte("Bad input,Input error,please reinput:\n"))

	case cmd == "get" && len(fields) == 2:
		fmt.Println("开始传送文件：", name)
		f, err := os.Open(name)
		if err != nil {
			log.Println(err)
			conn.Write([]byte("badinput"))
			return
		}
		defer f.Close()
		io.Copy(conn, f)
		fmt.Println("文件传送完毕")

	case cmd == "put" && len(fields) == 1:
		fmt.Println("Input error,please reinput:")
		conn.Write([]byte("Bad input,Input error,please reinput:\n"))

	case cmd == "put" && len(fields) == 2:
		//os.MkdirAll(filepath.Dir(name), 0755)
		f, err := os.Create(name)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(line)
		//fmt.Println(r)
		io.Copy(f, r)
		f.Close()
		conn.Close()

	case cmd == "exit":
		conn.Close()
	}
}

func main() {
	addr := "127.0.0.1:8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConn(conn)
	}
}
