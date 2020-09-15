package main

import (
	"calculadora/impl"
	"fmt"
	"net"
	"os"
	"shared"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

func CalculatorServerUDP() {
	msgFromClient := make([]byte, 1024)

	// Resolve server address
	addr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(shared.CALCULATOR_PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Listen on udp port
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("Server is ready (UDP)...")

	for {
		// Receive request
		n, addr, err := conn.ReadFromUDP(msgFromClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Handle request
		go HandleUDP(conn, msgFromClient, n, addr)
	}
}

func HandleUDP(conn *net.UDPConn, msgFromClient []byte, n int, addr *net.UDPAddr) {
	var msgToClient []byte
	var request shared.Request

	for {
		var json = jsoniter.ConfigCompatibleWithStandardLibrary

		//Deserialize request from clientserver
		err := json.Unmarshal(msgFromClient[:n], &request)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Process request
		r := impl.Calculadora{}.InvocaCalculadora(request)

		// Create response
		rep := shared.Reply{Result: []interface{}{r}}

		// Serialise response
		msgToClient, err = json.Marshal(rep)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Send response
		_, err = conn.WriteTo(msgToClient, addr)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Receive request
		n, addr, err = conn.ReadFromUDP(msgFromClient)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
	}
}

func main() {

	go CalculatorServerUDP()

	fmt.Scanln()
}
