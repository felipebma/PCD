#!/bin/bash

go run client.go & go run client.go & go run client_metrics.go 2> TCP_05.csv &  go run client.go & go run client.go
