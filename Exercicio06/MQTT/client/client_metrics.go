package main

import (
	"encoding/json"
	"fmt"
	"os"
	"shared"
	"strconv"
	"sync"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const qos = 1

var wg sync.WaitGroup

func main() {
	start := time.Now()

	// configurar cliente
	opts := MQTT.NewClientOptions()
	opts.AddBroker(shared.MQTTHost)
	clientID := shared.RandomString(32)
	opts.SetClientID(clientID)

	// criar novo cliente do broker
	client := MQTT.NewClient(opts)

	// conectar ao broker
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// desconectar cliente do broker
	defer client.Disconnect(250)

	// subscrever a um topico & usar um handler para receber as mensagens
	token = client.Subscribe(shared.MQTTReply+clientID, qos, receiveHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	wg.Add(shared.SampleSize)
	// loop
	for i := 0; i < shared.SampleSize; i++ {
		// cria a mensagem
		msg, err := json.Marshal(shared.Request{ClientID: clientID, Keywords: "Harry Potter"})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		//publicar a mensagem
		token := client.Publish(shared.MQTTRequest, qos, false, msg)
		token.Wait()
		if token.Error() != nil {
			fmt.Println(token.Error())
			os.Exit(1)
		}
	}
	wg.Wait()
	total := time.Now().Sub(start).Nanoseconds()
	fmt.Fprintf(os.Stderr, strconv.FormatInt(total, 10))
}

var receiveHandler MQTT.MessageHandler = func(c MQTT.Client, m MQTT.Message) {
	rep := shared.Reply{}
	err := json.Unmarshal(m.Payload(), &rep)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Books: [%s]\n", rep.Books)
	wg.Done()
}
