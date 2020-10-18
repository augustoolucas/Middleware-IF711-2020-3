package main

import (
	"encoding/json"
	"fmt"
	"github.com/gonum/stat"
	"github.com/streadway/amqp"
	"shared"
	"sync"
	"time"
)

func CalculatorClientRabbitMQ(clientID int, means *[]float64, stds *[]float64, wg *sync.WaitGroup, mtx *sync.Mutex) {
	// conecta ao server de mensageria
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao server de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o server de mensageria")
	defer ch.Close()
	defer wg.Done()

	// declara as filas
	requestQueue, err := ch.QueueDeclare(
		"request", false, false, false, false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no server de mensageria")

	replyQueue, err := ch.QueueDeclare(
		"response", false, false, false, false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no server de mensageria")

	// cria consumidor
	msgsFromServer, err := ch.Consume(replyQueue.Name, "", true, false,
		false, false, nil)
	shared.ChecaErro(err, "Falha ao registrar o consumidor server de mensageria")

	var responseTimes []float64
	var totalTime time.Duration

	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		start := time.Now()
		// t1 := time.Now()

		// prepara request
		msgRequest := shared.Request{Op: "add", P1: i, P2: i}
		msgRequestBytes, err := json.Marshal(msgRequest)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		// publica request
		err = ch.Publish("", requestQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o server de mensageria")

		// recebe resposta
		x := <-msgsFromServer

		var result shared.Reply
		json.Unmarshal(x.Body, &result)
		// fmt.Println(result.Result[0])

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
		go CalculatorClientRabbitMQ(i, &means, &stds, &wg, &mtx)
	}

	wg.Wait()

	meanFloat := stat.Mean(means, nil)
	stdDevFloat := stat.Mean(stds, nil)
	mean := time.Duration(meanFloat)
	stdDev := time.Duration(stdDevFloat)

	fmt.Println("Total execution time: ", time.Since(start), "- Average Mean: ", mean,
		" - Average Standard Deviation: ", stdDev)

}
