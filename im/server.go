package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	MapLock   sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

func (this *Server) start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()

	go this.ListenMessager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("connection success")
	user := NewUser(conn, this)

	isAlive := make(chan bool)

	user.Online()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Connection Read Error: ", err)
				return
			}
			msg := string(buf[:n-1])
			user.DoMessage(msg)
			isAlive <- true
		}
	}()
	for {
		select {
		case <-isAlive:
			// do nothing
		case <-time.After(time.Second * 60):
			user.SendMsg("you are offline!!!\n")
			close(user.C)
			conn.Close()
			return
		}
	}

}

func (this *Server) BoardCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg

	this.Message <- sendMsg

}

func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		this.MapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.MapLock.Unlock()
	}
}
