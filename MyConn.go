package main

import (
	"bufio"
	"fmt"
	"net"
)

type MyConn struct {
	Conn      *net.Conn
	Reader    *bufio.Reader
	Writer    *bufio.Writer
	InputChan chan []byte
}

func (c *MyConn) start() {
	go func() {
		for n := range c.InputChan {
			fmt.Println("input")
			fmt.Println(n)
			c.Writer.Write(n)
			c.Writer.Flush()
		}
	}()
}
