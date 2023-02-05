package server

import (
	"fmt"
	"net"
	"sync"
)

var (
	openConnections = make(map[string]Clients)
	Connection      = make(chan Message)
	deadConnection  = make(chan Message)
	joinConnection  = make(chan Message)
	history         []string
	mutex           = &sync.Mutex{}
)

type Clients struct {
	UserName string
	conn     net.Conn
}

type Message struct {
	msg      string
	time     string
	UserName string
}

func Server(port string) error {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("There is an issue in launching the server")
		return err
	}
	defer ln.Close()
	go Broadcaster()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("There is an issue in accepting the connection")
			return err
		}
		go HandleConnections(conn)
	}
}
