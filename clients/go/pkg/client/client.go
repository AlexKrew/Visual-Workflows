package client

import (
	"context"
	"log"
	"time"
	pb "workflows/gateway"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConfig struct {
	ServerAddress string
}

type Client struct {
	Client *pb.GatewayClient
	conn   *grpc.ClientConn
}

func StartClient(config ClientConfig) (*Client, error) {

	conn, err := grpc.Dial(config.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	gwClient := pb.NewGatewayClient(conn)

	return &Client{
		Client: &gwClient,
		conn:   conn,
	}, nil
}

func (client *Client) CheckHealth() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := (*client.Client).CheckHealth(ctx, &pb.Ping{Ping: 3})
	if err != nil {
		return err
	}

	log.Printf("Pong: %d", response.GetPong())

	return nil
}

func (client *Client) Close() {
	client.conn.Close()
}
