package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run client.go <host> <port>")
		os.Exit(1)
	}

	address := fmt.Sprintf("%v:%v", os.Args[1], os.Args[2])

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("connected to", address)

	serverReader := bufio.NewReader(conn)
	stdinReader := bufio.NewReader(os.Stdin)

	readWelcomeMsg(serverReader)

	for {
		// get msg length
		lenLine, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read from server:", err)
			return
		}

		size, _ := strconv.Atoi(strings.TrimSpace(lenLine))
		buf := make([]byte, size)
		io.ReadFull(serverReader, buf)
		fmt.Printf("%s\n", buf)

		// make move
		fmt.Print("> ")
		line, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read stdin:", err)
			return
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("failed to write to server:", err)
			return
		}
	}
}

func readWelcomeMsg(serverReader *bufio.Reader) {
	welcomeMsg, err := serverReader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read from server:", err)
		return
	}
	fmt.Print("lenLine: ", welcomeMsg)
}
