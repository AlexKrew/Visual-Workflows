package main

import (
	"workflows/clients/go/pkg/client"
)

var (
	// the connection address of the gateway grpc server
	CONN_ADDR = "localhost:50051"
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
