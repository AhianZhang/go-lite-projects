package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("connection remote error", err)
		return nil
	}
	client.conn = conn
	return client
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1.public chat")
	fmt.Println("2.private chat")
	fmt.Println("3.rename")
	fmt.Println("0.quit")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("number incorrect")
		return false
	}
}
func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {
		}
		switch client.flag {
		case 1:
			fmt.Println("public")
			client.PublicChat()
			break
		case 2:
			fmt.Println("private")
			client.PrivateChat()
			break
		case 3:
			fmt.Println("rename")
			client.Rename()
			break
		}

	}
}

func (client *Client) Rename() bool {
	fmt.Println(">>>input your new name")
	fmt.Scanln(&client.Name)

	sendMsg := "/rename " + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("rename failed", err)
		return false
	}
	return true
}
func (client *Client) PublicChat() {
	var chatMsg string
	fmt.Println(">>> input your message,[exit] to quit")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		chatMsg = ""
		fmt.Println(">>> input your message,[exit] to quit")
		fmt.Scanln(&chatMsg)
	}

}
func (client *Client) ListUsers() {
	msg := "/list\n"
	_, err := client.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	client.ListUsers()
	fmt.Println(">>> input username,[exit] to quit")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>> input your chat messages")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "/to," + remoteName + "," + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			chatMsg = ""
			fmt.Println(">>> input your chat messages")
			fmt.Scanln(&chatMsg)
		}
		client.ListUsers()
		fmt.Println(">>> input username,[exit] to quit")
		fmt.Scanln(&remoteName)
	}
}
func (client *Client) DealResp() {
	io.Copy(os.Stdout, client.conn)
	//equals below code
	//for {
	//	buf := make([]byte,4096)
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "server ip,")
	flag.IntVar(&serverPort, "port", 8888, "server port")
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("connect remote failed!")
		return
	}

	go client.DealResp()
	fmt.Println("connection remote success")
	client.Run()

}
