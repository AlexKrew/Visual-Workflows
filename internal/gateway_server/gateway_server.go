package gatewayserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	pb "workflows/gateway"
	"workflows/shared/shared_entities"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGatewayServer
}

type GatewayServer struct {
	activateJobStreams map[string][]*pb.Gateway_ActivateJobServer
	roundRobinIndex    map[string]int

	keepAliveChan chan any
}

var gatewayServer GatewayServer

func StartGatewayServer(port int) (*GatewayServer, error) {

	server := grpc.NewServer()
	pb.RegisterGatewayServer(server, &Server{})

	gatewayServer = GatewayServer{
		activateJobStreams: make(map[string][]*pb.Gateway_ActivateJobServer),
		roundRobinIndex:    make(map[string]int),
	}

	go startJobListener()
	go startServer(server, port)

	return &gatewayServer, nil
}

func startJobListener() error {
	// for {
	// 	// blocks until a new job is emitted by the channel
	// 	newJob := <-gatewayServer.jobQueue.NewJobs

	// 	streams, exists := gatewayServer.activateJobStreams[newJob.NodeType]
	// 	if !exists {
	// 		log.Printf("no worker client for jobtype '%s' registered", newJob.NodeType)
	// 		// TODO: Handle?
	// 		continue
	// 	}
	// 	index := gatewayServer.roundRobinIndex[newJob.NodeType]
	// 	stream := streams[index]

	// 	// if round-robin index is out-of-bound: reset to 0
	// 	if index+1 >= len(streams) {
	// 		gatewayServer.roundRobinIndex[newJob.NodeType] = 0
	// 	} else {
	// 		gatewayServer.roundRobinIndex[newJob.NodeType]++
	// 	}

	// 	jobInput, err := json.Marshal(newJob.Input)
	// 	if err != nil {
	// 		log.Fatalf("failed to transform jobinput into json: %s", err.Error())
	// 		continue
	// 	}

	// 	job := &pb.ActivatedJob{
	// 		JobId:      newJob.ID,
	// 		Type:       newJob.NodeType,
	// 		WorkflowId: newJob.WorkflowID,
	// 		Input:      string(jobInput),
	// 	}
	// 	(*stream).Send(&pb.ActivateJobResponse{
	// 		Job: job,
	// 	})
	// }
	return nil
}

func startServer(server *grpc.Server, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	log.Printf("gateway server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve gateway server: %v", err)
	}
}

func (gwServer *GatewayServer) keepAliveStream() {
	// does not complete until an item is inserted into the channel
	<-gwServer.keepAliveChan
}

func (gwServer *GatewayServer) CanExecute(jobType string) bool {
	_, hasClient := gwServer.activateJobStreams[jobType]
	return hasClient
}

func (gwServer *GatewayServer) Execute(job shared_entities.Job) error {
	streams, exists := gwServer.activateJobStreams[job.Type]
	if !exists {
		log.Printf("[GatewayServer]: stream index out-of-bound")
		return errors.New("no streams for jobtype")
	}

	nextStreamIndex := gwServer.roundRobinIndex[job.Type]
	if nextStreamIndex >= len(streams) {
		log.Printf("[GatewayServer]: stream index out-of-bound")
		return errors.New("stream index out of bound")
	}
	stream := streams[nextStreamIndex]

	jobInput, err := json.Marshal(job.Input)
	if err != nil {
		log.Printf("[GatewayServer]: failed to convert job to activatejob: %s", err.Error())
		return err
	}
	response := &pb.ActivateJobResponse{
		Job: &pb.ActivatedJob{
			JobId:      job.ID,
			Type:       job.Type,
			WorkflowId: job.WorkflowID,
			Input:      string(jobInput),
		},
	}
	err = (*stream).Send(response)
	if err != nil {
		log.Printf("failed to send ActivateJobResponse: %s", err.Error())
		gatewayServer.removeStream(job.Type, nextStreamIndex)
	} else {
		gatewayServer.increaseRoundRobin(job.Type)
	}

	return nil
}

func (gwServer *GatewayServer) removeStream(jobType string, index int) {
	// TODO:
}

func (gwServer *GatewayServer) decreaseRoundRobind(jobType string) {
	gwServer.roundRobinIndex[jobType]--

	if gwServer.roundRobinIndex[jobType] < 0 {
		gwServer.roundRobinIndex[jobType] = 0
	}
}

func (gwServer *GatewayServer) increaseRoundRobin(jobType string) {
	gwServer.roundRobinIndex[jobType]++

	if gwServer.roundRobinIndex[jobType] >= len(gwServer.activateJobStreams[jobType]) {
		gwServer.roundRobinIndex[jobType] = 0
	}
}

// grpc service implementations

func (server *Server) CheckHealth(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	log.Printf("Received: %v", in.GetPing())
	return &pb.Pong{Pong: in.GetPing() + 1}, nil
}

// ActivateJob opens a stream between the server and a gateway client.
// When there is a new job that need to be executed, the server can send it to the client
// through the stream. If the client completes the job, the grpc endpoint `CompleteJob` is addressed.
func (server *Server) ActivateJob(input *pb.ActivateJobRequest, stream pb.Gateway_ActivateJobServer) error {

	log.Println("Activate job connection established", input.Types)

	for _, jobtype := range input.GetTypes() {

		if _, ok := gatewayServer.activateJobStreams[jobtype]; !ok {
			gatewayServer.activateJobStreams[jobtype] = []*pb.Gateway_ActivateJobServer{}
			gatewayServer.roundRobinIndex[jobtype] = 0
		}
		gatewayServer.activateJobStreams[jobtype] = append(gatewayServer.activateJobStreams[jobtype], &stream)

		log.Printf("Added stream for job `%s`", jobtype)
	}

	// long-running stream
	gatewayServer.keepAliveStream()

	return nil
}

func (server *Server) CompleteJob(ctx context.Context, in *pb.CompleteJobRequest) (*pb.CompleteJobResponse, error) {
	return nil, errors.New("not implemented")
}
