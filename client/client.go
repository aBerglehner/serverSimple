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

	readServerInitMsg(serverReader)

	for {

		// data,err:=serverReader.ReadBytes('\n')

		lenLine, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read from server:", err)
			return
		}
		fmt.Print("lenLine: ", lenLine)

		size, _ := strconv.Atoi(strings.TrimSpace(lenLine))
		buf := make([]byte, size)
		io.ReadFull(serverReader, buf)
		fmt.Printf("%s\n", buf)

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

func readServerInitMsg(serverReader *bufio.Reader) string {
	welcome, err := serverReader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to read welcome:", err)
		return ""
	}
	fmt.Print("server: ", welcome)
	return welcome
}
