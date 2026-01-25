package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run client.go <host> <port>")
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%s", os.Args[1], os.Args[2])

	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("failed to connect:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("connected to", address)

	serverReader := bufio.NewReader(conn)
	stdinReader := bufio.NewReader(os.Stdin)

	readServerInitMsg(serverReader)

	for {
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

		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read from server:", err)
			return
		}

		fmt.Print("server: ", response)
	}
}

func readServerInitMsg(serverReader *bufio.Reader) {
	welcome, err := serverReader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read welcome:", err)
		return
	}
	fmt.Print("server: ", welcome)
}
