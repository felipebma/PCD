#!/bin/bash

go run client/client.go & go run client/client.go & go run client/client_metrics.go 2> MQTT_05.csv &  go run client/client.go & go run client/client.go
