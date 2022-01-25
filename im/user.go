package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	go user.ListenMessage()
	return user
}

func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

func (this *User) Online() {
	this.server.MapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.MapLock.Unlock()
	this.server.BoardCast(this, "online")
}
func (this *User) Offline() {
	this.server.MapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.MapLock.Unlock()
	this.server.BoardCast(this, "offline")
}
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}
func (this *User) DoMessage(msg string) {
	if msg == "/list" {

		this.server.MapLock.Lock()
		for _, u := range this.server.OnlineMap {
			listUser := "[" + u.Addr + "]" + u.Name + ": online!!! \n"
			this.SendMsg(listUser)
		}
		this.server.MapLock.Unlock()

	} else if len(msg) > 7 && msg[:7] == "/rename" {
		// /rename ahian
		newName := strings.Split(msg, " ")[1]
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("this name is tied up!!! ")
		} else {
			this.server.MapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.MapLock.Unlock()

			this.Name = newName
			this.SendMsg("update success,your new name is: " + newName + "\n")
		}
	} else if len(msg) > 3 && msg[:3] == "/to" {
		// /to,ahian,hello
		remoteName := strings.Split(msg, ",")[1]
		if remoteName == "" {
			this.SendMsg("please use: \"/to,username,message\"")
			return
		}
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("cannot find user")
			return
		}
		content := strings.Split(msg, ",")[2]
		if content == "" {
			this.SendMsg("message is empty")
			return
		}
		remoteUser.SendMsg(this.Name + " said to you : " + content + "\n")

	} else {

		this.server.BoardCast(this, msg)
	}
}
