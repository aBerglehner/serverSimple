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

	ch := make(chan string)
	curPlayer := 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection, err:", err)
			continue
		}
		curPlayer++
		if curPlayer == 1 {
			fmt.Printf("Player: %v joined\n", curPlayer)
			go connPlayer1(conn, ch)
		}
		if curPlayer == 2 {
			fmt.Printf("Player: %v joined\n", curPlayer)
			// TODO:
			go connPlayer2(conn, ch)
		}
		if curPlayer > 2 {
			// todo
			fmt.Println("!!! todo !!!")
		}
	}
}

func connPlayer2(conn net.Conn, ch chan string) {
	defer conn.Close()
	ch <- "halli galli"
	for {
		// todo
	}
}

func connPlayer1(conn net.Conn, ch chan string) {
	defer conn.Close()

	// greeting msg
	welcomeMsg := fmt.Sprintln("Waiting for another player!")
	_, err := conn.Write([]byte(welcomeMsg))
	if err != nil {
		fmt.Println("failed to write welcome message:", err)
		return
	}
	// todo if player 2 swap turn and set it to player 1
	// todo turn
	select {
	case msg := <-ch:
		fmt.Println("player 2 joined: ", msg)
	}

	// todo if player 2 start

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
