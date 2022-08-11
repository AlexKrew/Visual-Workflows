package client

import (
	"log"
	"workflows/clients/go/pkg/nodes"
	"workflows/clients/go/pkg/workers"
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

	manager *workers.JobManager
}

func StartClient(config ClientConfig) (*Client, error) {

	conn, err := grpc.Dial(config.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	gwClient := pb.NewGatewayClient(conn)

	manager := workers.NewJobManager()

	return &Client{
		Client:  &gwClient,
		conn:    conn,
		manager: &manager,
	}, nil
}

func (client *Client) AddJobWorker(jobType string, handler nodes.HandlerFunc) {
	worker := workers.JobWorker{
		JobType: jobType,
		Handler: handler,
	}

	client.manager.AddWorker(worker)
}

func (client *Client) StartJobPolling() {
	log.Println("Start polling ...")
}

func (client *Client) Close() {
	client.conn.Close()
}
