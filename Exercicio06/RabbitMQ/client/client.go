package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"shared"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	// gera nova seed
	rand.Seed(time.Now().UTC().UnixNano())

	// conecta ao broker
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	shared.ChecaErro(err, "Não foi possível se conectar ao servidor de mensageria")
	defer conn.Close()

	// cria o canal
	ch, err := conn.Channel()
	shared.ChecaErro(err, "Não foi possível estabelecer um canal de comunicação com o servidor de mensageria")
	defer ch.Close()

	queue_name := shared.RandomString(32)

	// declara a fila para as respostas
	replyQueue, err := ch.QueueDeclare(
		queue_name, // name
		false,      // durable
		false,      // delete when unused
		true,       // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	// cria servidor da fila de response
	msgs, err := ch.Consume(
		replyQueue.Name, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	shared.ChecaErro(err, "Falha ao registrar o servidor no broker")

	for i := 0; i < shared.SampleSize; i++ {
		// prepara mensagem
		msgRequest := shared.Request{Keywords: "Harry Potter"}
		msgRequestBytes, err := json.Marshal(msgRequest)
		shared.ChecaErro(err, "Falha ao serializar a mensagem")

		correlationID := shared.RandomString(32)

		err = ch.Publish(
			"",                  // exchange
			shared.RequestQueue, // routing key
			false,               // mandatory
			false,               // immediate
			amqp.Publishing{ // publishing
				ContentType:   "text/plain",    // content type
				CorrelationId: correlationID,   // correlation id
				ReplyTo:       queue_name,      // reply queue
				Body:          msgRequestBytes, // body
			},
		)

		// recebe mensagem do servidor de mensageria
		m := <-msgs

		// deserializada e imprime mensagem na tela
		msgResponse := shared.Reply{}
		err = json.Unmarshal(m.Body, &msgResponse)
		shared.ChecaErro(err, "Erro na deserialização da resposta")
		fmt.Printf("Books: [%s]\n\n", msgResponse.Books)
	}
}
