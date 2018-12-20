package main

import (
	"bufio"
	"net"
)

type MyConn struct {
	Conn *net.Conn
	Reader *bufio.Reader
	Writer *bufio.Writer
	InputChan chan []byte
}

func (c *MyConn)start(){
	go func() {
		for n := range c.InputChan {
			c.Writer.Write(n)
		}
	}()
}
