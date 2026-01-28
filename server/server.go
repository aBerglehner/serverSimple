package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/alex/serverSimple/server/board"
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

	startCh := make(chan string)
	board := board.NewBoard()
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
			go connPlayer1(conn, startCh, board)
		}
		if curPlayer == 2 {
			fmt.Printf("Player: %v joined\n", curPlayer)
			// TODO:
			go connPlayer2(conn, startCh, board)
		}
		if curPlayer > 2 {
			// todo
			fmt.Println("!!! todo !!!")
		}
	}
}

func connPlayer2(conn net.Conn, startCh chan<- string, board board.Board) {
	defer conn.Close()

	startCh <- "start"
	for {
		// todo
		time.Sleep(1 * time.Second)
	}
}

func connPlayer1(conn net.Conn, startCh <-chan string, board board.Board) {
	defer conn.Close()

	// greeting msg
	welcomeMsg := fmt.Sprintln("Waiting for another player!")
	_, err := conn.Write([]byte(welcomeMsg))
	if err != nil {
		fmt.Println("failed to write welcome message:", err)
		return
	}

	// waiting for 2 player
	<-startCh
	fmt.Println("after start")

	// start board
	payload := board.String()
	payload = fmt.Sprintf("%d\n%s", len(payload), payload)
	_, err = conn.Write([]byte(payload))
	if err != nil {
		fmt.Println("failed to write welcome message:", err)
		return
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
