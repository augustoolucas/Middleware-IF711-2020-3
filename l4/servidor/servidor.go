package main

import (
	"calculadora/impl"
	"encoding/json"
	"fmt"
	"log"
	"shared"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao server de mensageria")
	defer conn.Close()

	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o server de mensageria")
	defer ch.Close()

	// declaração de filas
	requestQueue, err := ch.QueueDeclare("request", false, false, false,
		false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no server de mensageria")

	replyQueue, err := ch.QueueDeclare("response", false, false, false,
		false, nil)
	shared.ChecaErro(err, "Não foi possível criar a fila no server de mensageria")

	// prepara o recebimento de mensagens do clientserver
	msgsFromClient, err := ch.Consume(requestQueue.Name, "", true, false,
		false, false, nil)
	shared.ChecaErro(err, "Falha ao registrar o consumidor do server de mensageria")

	fmt.Println("Servidor pronto...")
	for d := range msgsFromClient {

		// recebe request
		msgRequest := shared.Request{}
		err := json.Unmarshal(d.Body, &msgRequest)
		shared.ChecaErro(err, "Falha ao desserializar a mensagem")

		// processa request
		r := impl.Calculadora{}.InvocaCalculadora(msgRequest)

		// prepara resposta
		replyMsg := shared.Reply{Result: []interface{}{r}}
		replyMsgBytes, err := json.Marshal(replyMsg)
		shared.ChecaErro(err, "Não foi possível criar a fila no server de mensageria")
		if err != nil {
			log.Fatalf("%s: %s", "Falha ao serializar mensagem", err)
		}

		// publica resposta
		err = ch.Publish("", replyQueue.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(replyMsgBytes)})
		shared.ChecaErro(err, "Falha ao enviar a mensagem para o server de mensageria")
	}
}
