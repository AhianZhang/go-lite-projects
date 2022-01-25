package main

import "fmt"

func main() {
	fmt.Println("111")
	server := NewServer("127.0.0.1", 8888)
	server.start()
}
