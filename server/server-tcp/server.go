package main

import (
	"calculadora/impl"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shared"
	"strconv"
)

func CalculatorServerTCP() {

	// Listen on tcp port
	ln, err := net.Listen("tcp", "localhost:"+strconv.Itoa(2308))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Server is ready to accept connections (TCP)...")

	for {
		// Accept new connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Handle connection
		go HandleTCP(conn)
	}
}

func HandleTCP(conn net.Conn) {
	var msgFromClient shared.Request

	// Close connection
	defer conn.Close()

	// Create coder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for {
		// Receive request
		err := jsonDecoder.Decode(&msgFromClient)
		if err != nil && err.Error() == "EOF" {
			conn.Close()
			break
		}

		// Process request
		r := impl.Calculadora{}.InvocaCalculadora(msgFromClient)

		// Create response
		msgToClient := shared.Reply{[]interface{}{r}}

		// Serialise and send response to clientserver
		err = jsonEncoder.Encode(msgToClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func main() {

	go CalculatorServerTCP()

	fmt.Scanln()
}
