package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shared"
	"strconv"
	"sync"
	"time"

	"github.com/gonum/stat"
)

func CalculatorClientUDP(clientID int, wg *sync.WaitGroup, mux *sync.Mutex) {
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
	defer wg.Done()

	var responseTimes []float64
	var totalTime time.Duration

	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		start := time.Now()
		// Create request
		request := shared.Request{Op: "add", P1: i, P2: i}

		//mux.Lock()
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

		executionTime := time.Since(start)
		//mux.Unlock()
		responseTimes = append(responseTimes, float64(executionTime))
		totalTime += executionTime
		// fmt.Printf("%s(%d,%d) = %.0f \n", request.Op, request.P1, request.P2, response.Result[0])

		//time.Sleep(time.Second * 2)
	}

	meanFloat, stdDevFloat := stat.MeanStdDev(responseTimes, nil)
	mean := time.Duration(meanFloat)
	stdDev := time.Duration(stdDevFloat)
	fmt.Println("ID: ", clientID, "- Total Time: ", totalTime, "- Mean: ",
		mean, " - Standard Deviation: ", stdDev)
}

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex

	start := time.Now()

	for i := 0; i < shared.CLIENTS; i++ {
		wg.Add(1)
		go CalculatorClientUDP(i, &wg, &mux)
	}

	wg.Wait()

	fmt.Println("Total execution time: ", time.Since(start))

	_, _ = fmt.Scanln()
}
