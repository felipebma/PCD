#!/bin/bash

go run client.go & go run client.go & go run client.go & go run client.go & go run client_metrics.go 2> UDP_10.csv & \
go run client.go & go run client.go & go run client.go & go run client.go & go run client.go
