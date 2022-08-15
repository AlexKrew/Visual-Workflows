package main

import (
	"time"
	"workflows/clients/go/pkg/client"
)

var (
	// the connection address of the gateway grpc server
	CONN_ADDR = "localhost:50051"
	// the keep-alive duration for the request streams
	POLL_DURATION = 30 * time.Second
)

func main() {
	clientConfig := client.ClientConfig{
		ServerAddress: CONN_ADDR,
	}
	gatewayClient, err := client.StartClient(clientConfig)
	if err != nil {
		panic(err)
	}

	gatewayClient.StartJobPolling()
}
