package main

import (
	"calculadora/grpc/calculadora"
	"fmt"
	"shared"
	"strconv"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var idx int32

	// Estabelece conexão com o server
	conn, err := grpc.Dial("localhost"+":"+strconv.Itoa(shared.CALCULATOR_PORT), grpc.WithInsecure())
	shared.ChecaErro(err, "Não foi possível se conectar ao server")
	defer conn.Close()

	calc := calculadora.NewCalculatorClient(conn)

	// contacta o server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for idx = 0; idx < shared.SAMPLE_SIZE; idx++ {

		t1 := time.Now()
		// invoca operação remota
		rep, msgErr := calc.Add(ctx, &calculadora.Request{Op: "add", P1: idx, P2: idx})

		if msgErr != nil {
			fmt.Println(idx, msgErr)
		} else {
			x := float64(time.Since(t1))

			fmt.Printf("%f %d\n", x, rep.N)
		}
	}
}
