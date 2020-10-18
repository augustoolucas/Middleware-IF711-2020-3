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

func CalculatorClientRPC(clientID int, means *[]float64, stds *[]float64, wg *sync.WaitGroup, mtx *sync.Mutex) {
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
		go CalculatorClientRPC(i, &means, &stds, &wg, &mtx)
	}

	wg.Wait()

	meanFloat := stat.Mean(means, nil)
	stdDevFloat := stat.Mean(stds, nil)
	mean := time.Duration(meanFloat)
	stdDev := time.Duration(stdDevFloat)

	fmt.Println("Total execution time: ", time.Since(start), "- Average Mean: ", mean,
		" - Average Standard Deviation: ", stdDev)
}
