package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shared"
	"strconv"
)

func CalculatorClientTCP() {
	var response shared.Reply

	// Connect to server
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(shared.CALCULATOR_PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Close connection
	defer conn.Close()

	// Create enconder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		// Prepare request
		msgToServer := shared.Request{"add", i, i}

		// Serialise and send request to server
		err = jsonEncoder.Encode(msgToServer)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		// Receive response from server
		err = jsonDecoder.Decode(&response)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		fmt.Printf("%s(%d,%d) = %.0f \n",msgToServer.Op,msgToServer.P1,msgToServer.P2,response.Result[0].(float64))
	}
}

func main() {

	go CalculatorClientTCP()

	_, _ = fmt.Scanln()
}

