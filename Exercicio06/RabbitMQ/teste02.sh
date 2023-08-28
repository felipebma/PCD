#!/bin/bash

go run client/client.go & go run client/client_metrics.go 2> RabbitMQ_02.csv 
