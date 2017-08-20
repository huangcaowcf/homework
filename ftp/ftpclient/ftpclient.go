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

func main() {
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:8021")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("FTP服务器连接成功，请输入ls|get $filename|put $filename:")
		defer conn.Close()
		fmt.Print("ftp > ")

		input := bufio.NewReader(os.Stdin)
		line, err := input.ReadString('\n')
		if err != nil {
			break
		}
		//fmt.Println("line:", line)
		if line == "\n" {
			conn.Close()
			continue
		}
		var cmd string
		var name string
		fields := strings.Fields(line)
		if len(fields) == 1 {
			cmd = fields[0]
		} else if len(fields) == 2 {
			cmd = fields[0]
			name = fields[1]
		} else {
			fmt.Println("Input error,please reinput:")
			continue
		}

		//r := bufio.NewReader(conn)
		//fmt.Println(line)
		defer conn.Close()

		switch cmd {
		case "ls":
			conn.Write([]byte(line))
			io.Copy(os.Stdout, conn)

		case "pwd":
			conn.Write([]byte(line))
			io.Copy(os.Stdout, conn)

		case "lpwd":
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(dir)
			conn.Close()

		case "lls":
			dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println("Ls本当前目录:", dir)
			info, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range info {
				fmt.Println(file.Name())
			}
			conn.Close()

		case "get":
			fmt.Println("开始下载文件")
			//os.MkdirAll(filepath.Dir(name), 0755)
			file, err1 := os.Open(name)

			if err1 != nil {
				conn.Write([]byte(line))
				r := bufio.NewReader(conn)
				lines, _ := r.ReadString('\n')
				lines = strings.TrimSpace(lines)
				fields := strings.Fields(lines)
				cmd := fields[0]
				//fmt.Println(cmd)
				if cmd == "badinput" {
					fmt.Println("服务器没有这个文件，请重新输入：")
					conn.Close()
				} else {
					f, err2 := os.Create(name)
					if err2 != nil {
						log.Print(err2)
						continue
					}
					io.Copy(f, r)
					fmt.Println("文件下载完毕")
					f.Close()
				}
			} else {
				fmt.Println("文件已存在，请重新输入：")
				conn.Close()
				continue
			}
			file.Close()

		case "put":
			fmt.Println("开始上传文件")
			f, err := os.Open(name)
			if err != nil {
				log.Print("没有这个文件，请重新输入：")
				conn.Close()
				//conn.Write([]byte("bad"))
			} else {
				conn.Write([]byte(line))
				fmt.Println("文件上传完毕")
				io.Copy(conn, f)
				f.Close()
				conn.Close()
			}

		case "help", "Help", "HELP":

			fmt.Println("usage:")
			fmt.Println("ls           list server's files")
			fmt.Println("lls          list client's files")
			fmt.Println("pwd          print server's current dir")
			fmt.Println("lpwd         print client's current dir")
			fmt.Println("get [file]   download file")
			fmt.Println("put [file]   upload file")
			fmt.Println("exit/exit()  exit")
			fmt.Println("help/Help/HELP   print the help info")
			conn.Close()
		case "exit", "exit()", "Exit":
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("Input error,please reinput:")
			conn.Close()
			continue
		}

	}
}
