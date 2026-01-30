package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"

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

	player2Turn := make(chan struct{})
	player1Turn := make(chan struct{})

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
			go connPlayer1(conn, player1Turn, player2Turn, board)
		}
		if curPlayer == 2 {
			fmt.Printf("Player: %v joined\n", curPlayer)
			go connPlayer2(conn, player2Turn, player1Turn, board)
		}
		if curPlayer > 2 {
			// todo
			fmt.Println("!!! todo !!!")
		}
	}
}

func connPlayer2(conn net.Conn, ownTurn chan<- struct{}, player1Turn <-chan struct{}, myBoard *board.Board) {
	defer conn.Close()
	ownTurn <- struct{}{}

	reader := bufio.NewReader(conn)
	for {
		// waiting for player 1
		<-player1Turn
		// fmt.Println("after player 1 turn")

		payload := myBoard.String()
		payload = fmt.Sprintf("%d\n%s", len(payload), payload)
		_, err := conn.Write([]byte(payload))
		if err != nil {
			fmt.Println("failed to write welcome message:", err)
			return
		}

		requestMsg, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request player 2: %s\n", requestMsg)

		myBoard.Update(board.O, string(requestMsg))
		ownTurn <- struct{}{}
	}
}

func connPlayer1(conn net.Conn, ownTurn chan<- struct{}, player2Turn <-chan struct{}, myBoard *board.Board) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	// get board -> make move -> get board -> make move
	for {
		// waiting for player 2
		<-player2Turn
		// fmt.Println("after player 2 turn")

		payload := myBoard.String()
		payload = fmt.Sprintf("%d\n%s", len(payload), payload)
		_, err := conn.Write([]byte(payload))
		if err != nil {
			fmt.Println("failed to write welcome message:", err)
			return
		}

		requestMsg, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}
		fmt.Printf("request player 1: %s\n", requestMsg)

		myBoard.Update(board.X, string(requestMsg))
		ownTurn <- struct{}{}
	}
}
