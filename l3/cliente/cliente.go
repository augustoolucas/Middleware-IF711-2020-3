package main

import (
	"calculadora/grpc/calculadora"
	"fmt"
	"github.com/gonum/stat"
	"shared"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func CalculatorClientRPC(clientID int, wg *sync.WaitGroup) {
	var idx int32

	// Estabelece conexão com o server
	conn, err := grpc.Dial("localhost"+":"+strconv.Itoa(shared.CALCULATOR_PORT), grpc.WithInsecure())
	shared.ChecaErro(err, "Não foi possível se conectar ao server")
	defer conn.Close()

	calc := calculadora.NewCalculatorClient(conn)

	// contacta o server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	defer wg.Done()

	var responseTimes []float64
	var totalTime time.Duration

	for idx = 0; idx < shared.SAMPLE_SIZE; idx++ {
		start := time.Now()

		// invoca operação remota
		_, msgErr := calc.Add(ctx, &calculadora.Request{Op: "add", P1: idx, P2: idx})

		if msgErr != nil {
			fmt.Println(idx, msgErr)
		}

		//fmt.Printf("%f %d\n", x, rep.N)

		executionTime := time.Since(start)
		responseTimes = append(responseTimes, float64(executionTime))
		totalTime += executionTime
	}

	meanFloat, stdDevFloat := stat.MeanStdDev(responseTimes, nil)
	mean := time.Duration(meanFloat)
	stdDev := time.Duration(stdDevFloat)
	fmt.Println("ID: ", clientID, "Total time: ", totalTime, "- Mean: ", mean,
		" - Standard Deviation: ", stdDev)
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < shared.CLIENTS; i++ {
		wg.Add(1)
		go CalculatorClientRPC(i, &wg)
	}

	wg.Wait()

	fmt.Println("Total execution time: ", time.Since(start))
}
