package main

import (
	"encoding/json"
	"fmt"
	"github.com/gonum/stat"
	"net"
	"os"
	"shared"
	"strconv"
	"sync"
	"time"
)

func CalculatorClientTCP(clientID int, means *[]float64, stds *[]float64, wg *sync.WaitGroup, mtx *sync.Mutex) {
	var response shared.Reply

	// Connect to server
	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(shared.CALCULATOR_PORT))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Close connection
	defer conn.Close()
	defer wg.Done()

	// Create enconder/decoder
	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	var responseTimes []float64
	var totalTime time.Duration

	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		start := time.Now()
		// Prepare request
		msgToServer := shared.Request{Op: "add", P1: i, P2: i}

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

		executionTime := time.Since(start)
		responseTimes = append(responseTimes, float64(executionTime))
		totalTime += executionTime

		//fmt.Printf("ID: %d - %s(%d,%d) = %.0f \n", clientID, msgToServer.Op,
		//	msgToServer.P1, msgToServer.P2, response.Result[0].(float64))
	}

	meanFloat, stdDevFloat := stat.MeanStdDev(responseTimes, nil)
	mean := time.Duration(meanFloat)
	stdDev := time.Duration(stdDevFloat)
	fmt.Println("ID: ", clientID, "- Total time: ", totalTime, "- Mean: ",
		mean, " - Standard Deviation: ", stdDev)

	mtx.Lock()
	*means = append(*means, meanFloat)
	*stds = append(*stds, stdDevFloat)
	mtx.Unlock()
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	var means []float64
	var stds []float64
	var mtx sync.Mutex

	for i := 0; i < shared.CLIENTS; i++ {
		wg.Add(1)
		go CalculatorClientTCP(i, &means, &stds, &wg, &mtx)
	}

	wg.Wait()

	meanFloat := stat.Mean(means, nil)
	stdDevFloat := stat.Mean(stds, nil)
	mean := time.Duration(meanFloat)
	stdDev := time.Duration(stdDevFloat)

	fmt.Println("Total execution time: ", time.Since(start), "- Average Mean: ",
		mean, " - Average Standard Deviation: ", stdDev)
}
