package main

import (
	"bufio"
	"net"
	"sync"
)

type MyServer struct {
	Conns map[string]*MyConn
	RwLock sync.RWMutex
}
func (s * MyServer) start(){
	ln,_:=net.Listen("tcp","127.0.0.1:3003")
	defer ln.Close()
	for {
		conn,_ := ln.Accept()
		reader:= bufio.NewReader(conn);
		first,_ := reader.ReadString('\n')
		writer:=bufio.NewWriter(conn)
		myConn:=MyConn{
			Conn:&conn,
			Reader:reader,
			Writer:writer,
			InputChan:make(chan []byte),
		}
		myConn.start()
		s.RwLock.Lock()
		s.Conns[first]=&myConn
		s.RwLock.Unlock()
	}
}
func main() {
	
}
