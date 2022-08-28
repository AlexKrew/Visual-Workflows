package client

import (
	"context"
	"io"
	"log"
	"time"
	"workflows/example_services/random_service"
	pb "workflows/gateway"
	"workflows/shared/job_manager"
	"workflows/shared/job_worker"
	"workflows/shared/shared_entities"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientConfig struct {
	ServerAddress string
}

type Client struct {
	Client *pb.GatewayClient
	conn   *grpc.ClientConn

	manager *job_manager.JobManager
}

func StartClient(config ClientConfig) (*Client, error) {

	conn, err := grpc.Dial(config.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panicf("initialization failed: %s", err.Error())
		return nil, err
	}

	gwClient := pb.NewGatewayClient(conn)

	manager := job_manager.NewJobManager()

	// injectWorker := job_worker.NewNodeJobWorker("Inject", nodes.ProcessInject)

	randomService := random_service.NewRandomService()
	randomServiceWorker := job_worker.NewServiceJobWorker("RandomService", randomService.DoRandomThingsAdapter)

	// manager.AddWorker(injectWorker)
	manager.AddWorker(randomServiceWorker)

	return &Client{
		Client:  &gwClient,
		conn:    conn,
		manager: &manager,
	}, nil
}

func (client *Client) StartJobPolling() {
	log.Println("Start polling ...")

	request := &pb.ActivateJobRequest{
		Types:   client.manager.SupportedServices(),
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
		go client.handleJob(response)
	}
}

func (client *Client) handleJob(activateJob *pb.ActivateJobResponse) {
	jobItem := activateJob.GetJob()

	job, err := shared_entities.JobFromJSONString(jobItem.Input)
	if err != nil {
		log.Println("Failed to extract job from json:", err)
		return
	}

	result, err := client.manager.Execute(job)
	if err != nil {
		log.Println("Failed to execute job:", err)
		return
	}

	outputJSON, err := result.ToJSONString()
	if err != nil {
		log.Println("Failed to convert output to json string", err)
		return
	}

	completeJob := &pb.CompleteJobRequest{
		JobId:      job.ID,
		WorkflowId: jobItem.GetWorkflowId(),
		Output:     outputJSON,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = (*client.Client).CompleteJob(ctx, completeJob)
	if err != nil {
		log.Println("Complete job failed:", err)
		return
	}
}

func (client *Client) Close() {
	client.conn.Close()
}
