package client

import (
	"context"
	"io"
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

	manager.AddWorker(workers.NewInjectWorker())

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

	request := &pb.ActivateJobRequest{
		Types:   client.manager.Types(),
		Timeout: 0,
	}
	stream, err := (*client.Client).ActivateJob(context.Background(), request)
	if err != nil {
		log.Printf("failed to listen for activate jobs: %s", err)
		return
	}

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			log.Println("EOF")
			break
		}

		if err != nil {
			log.Printf("failed to receive jobs: %s", err.Error())
			break
		}

		log.Printf("Received job to execute: %s", response)
	}
}

func (client *Client) Close() {
	client.conn.Close()
}
