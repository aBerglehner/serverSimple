package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}
	port := fmt.Sprintf(":%v", os.Args[1])
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to create listener, err: ", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("listening on: %v\n", listener.Addr())

	curPlayer := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		curPlayer++
		go handleConn(conn, curPlayer)
	}
}

func handleConn(conn net.Conn, player int) {
	defer func() {
		conn.Close()
		// todo if 1 player leaves the other wins
	}()
	fmt.Printf("Player: %v joined\n", player)

	// Send greeting immediately
	welcomeMsg := fmt.Sprintf("Welcome to the echo server player: %v\n", player)
	_, err := conn.Write([]byte(welcomeMsg))
	if err != nil {
		fmt.Println("failed to write welcome message:", err)
		return
	}

	reader := bufio.NewReader(conn)
	for {
		requestMsg, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request: %s\n", requestMsg)

		line := fmt.Sprintf("Echo: %s", requestMsg)
		// fmt.Printf("response: %v\n", line)

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("failed to write data, err:", err)
			return
		}
	}
}
