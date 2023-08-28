package main

import (
	"encoding/json"
	"fmt"
	"impl"
	"os"
	"shared"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const qos = 0

var count = 0

func main() {

	// configurar cliente
	opts := MQTT.NewClientOptions()
	opts.AddBroker(shared.MQTTHost)
	opts.SetClientID("subscriber 1")
	//opts.DefaultPublishHandler = receiveHandler

	// criar novo cliente
	client := MQTT.NewClient(opts)

	// conectar ao broker
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// desconectar ao broker
	defer client.Disconnect(250)

	// subscrever a um topico & usar um handler para receber as mensagens
	token = client.Subscribe(shared.MQTTRequest, qos, receiveHandler)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	fmt.Println("Consumidor iniciado...")
	fmt.Scanln()
}

var receiveHandler MQTT.MessageHandler = func(c MQTT.Client, m MQTT.Message) {
	go handleConnection(c, m)
}

func handleConnection(c MQTT.Client, m MQTT.Message) {
	count++
	req := shared.Request{}
	err := json.Unmarshal(m.Payload(), &req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := impl.Bookstore{}.FindBooks(req)
	rep, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.Marshal(rep)
	token := c.Publish(shared.MQTTReply+req.ClientID, qos, false, rep)
	token.Wait()
	if token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	println(count)
	// fmt.Printf("Recebida: ´%s´ Enviada: ´%s´\n", req, rep)

}
