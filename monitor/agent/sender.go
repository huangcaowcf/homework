package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/51reboot/golang-01-homework/lesson12/wuchuanfang/monitor/common"
)

type Sender struct {
	addr string
	ch   chan *common.Metric
}

func NewSender(addr string) *Sender {
	return &Sender{
		addr: addr,
		ch:   make(chan *common.Metric),
	}
}

/*
func (s *Sender) connect() net.Conn {
	n := 100 * time.Millisecond
	for {
		conn, err := net.Dial("tcp", s.addr)
		if err != nil {
			log.Println(err)
			n = n * 2
			if n > time.Second {

			}
		}
	}
}
*/
func (s *Sender) Start() {

	conn, err := net.Dial("tcp", s.addr)
	if err != nil {
		log.Println(err)
	}

	for {
		for metric := range s.ch {
			buf, _ := json.Marshal(metric)
			_, err := fmt.Fprintf(conn, "%s\n", buf)
			fmt.Println(string(buf))
			conn.Write(buf)
			if err != nil {
				conn.Close()
				//conn = s.connect()
				//log.Println(conn.LocalAddr())
			}

		}
	}
}

func (s *Sender) Channel() chan *common.Metric {
	return s.ch
}
