package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":4545")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}
}

// обработка подключения
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		input := make([]byte, (1024 * 4))
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}

		source := string(input[0:n])

		messageToClient := ""

		target, err := strconv.Atoi(source)
		if err != nil {
			messageToClient = strings.ToUpper(source)
		} else {
			messageToClient = strconv.Itoa(target * 2)
		}

		fmt.Println(source, "-", target)

		conn.Write([]byte(messageToClient))
	}
}
