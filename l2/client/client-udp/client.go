package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shared"
	"strconv"
	"time"
)

func CalculatorClientUDP() {
	var response shared.Reply

	// Resolve server address
	addr, err := net.ResolveUDPAddr("udp", "localhost:"+strconv.Itoa(shared.CALCULATOR_PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Connect to server
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	decoder := json.NewDecoder(conn)
	encoder := json.NewEncoder(conn)

	// Close connection
	defer conn.Close()

	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		// Create request
		request := shared.Request{Op: "add", P1: i, P2: i}

		// Serialise and send request
		err = encoder.Encode(request)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Receive response from server
		err = decoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Printf("%s(%d,%d) = %.0f \n", request.Op, request.P1, request.P2, response.Result[0])

		time.Sleep(time.Second * 2)
	}
}

func main() {

	go CalculatorClientUDP()

	_, _ = fmt.Scanln()
}

