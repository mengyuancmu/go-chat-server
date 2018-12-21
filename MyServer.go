package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type MyServer struct {
	Conns  map[string]*MyConn
	RwLock sync.RWMutex
}

func (s *MyServer) start() {
	ln, _ := net.Listen("tcp", "127.0.0.1:3003")
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		reader := bufio.NewReader(conn)
		first, _ := reader.ReadString('\n')
		fmt.Println(first)
		writer := bufio.NewWriter(conn)
		myConn := &MyConn{
			Conn:      &conn,
			Reader:    reader,
			Writer:    writer,
			InputChan: make(chan []byte),
		}
		myConn.start()
		s.RwLock.Lock()
		s.Conns[first] = myConn
		s.RwLock.Unlock()
		go func() {
			for {
				msg, _ := myConn.Reader.ReadString('\n')
				msgArr := strings.Split(msg, "::")
				if len(msgArr) == 2 {
					id := msgArr[0]
					msgBody := msgArr[1]
					s.RwLock.RLock()
					target := s.Conns[id]
					if target != nil {
						target.InputChan <- []byte(msgBody)
					}
					s.RwLock.RUnlock()
				}
			}
		}()
	}
}
func main() {
	server := &MyServer{
		Conns: make(map[string]*MyConn),
	}
	server.start()
}
