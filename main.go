package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("couldn't read file", err)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error while accepting")
			continue
		}
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read request line
	line, err := reader.ReadString('\n')


	if err != nil {
		fmt.Println("Read error:", err)
		return
	}


	// this is the same thing which we did last time with reading 8bit file , we are just splitting the request into 3 parts
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)
	if len(parts) < 3 {
		fmt.Println("Invalid request line:", line)
		return
	}
	method, path, proto := parts[0], parts[1], parts[2]
	fmt.Printf("Request: %s %s %s\n", method, path, proto)

	// Read headers
	headers := map[string]string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Header read error:", err)
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break // end of headers
		}
		kv := strings.SplitN(line, ":", 2)
		if len(kv) == 2 {
			headers[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		}
	}

	// Build and write response
	body := "Hello, world!"
	resp := fmt.Sprintf(
		"%s 200 OK\r\nContent-Length: %d\r\nContent-Type: text/plain\r\nConnection: close\r\n\r\n%s",
		proto, len(body), body,
	)

	conn.Write([]byte(resp))
	fmt.Println("Response sent")
}
