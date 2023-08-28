#!/bin/bash

go run client/client_metrics.go 2> MQTT_02.csv & go run client/client.go
