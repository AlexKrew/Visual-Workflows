package gatewayserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	pb "workflows/gateway"
	"workflows/internal/processors/workflow_processor"

	"github.com/goccy/go-json"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGatewayServer
}

type GatewayServer struct {
	activateJobStreams map[string][]*pb.Gateway_ActivateJobServer
	roundRobinIndex    map[string]int
	jobQueue           *workflow_processor.JobQueue
}

var gatewayServer GatewayServer

func StartGatewayServer(port int, jobQueue *workflow_processor.JobQueue) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterGatewayServer(server, &Server{})

	gatewayServer = GatewayServer{
		activateJobStreams: make(map[string][]*pb.Gateway_ActivateJobServer),
		roundRobinIndex:    make(map[string]int),
		jobQueue:           jobQueue,
	}

	go startJobListener()

	log.Printf("gateway server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve gateway server: %v", err)
		return err
	}

	return nil
}

func startJobListener() error {
	for {
		// blocks until a new job is emitted by the channel
		newJob := <-gatewayServer.jobQueue.NewJobs

		streams, exists := gatewayServer.activateJobStreams[newJob.NodeType]
		if !exists {
			log.Printf("no worker client for jobtype '%s' registered", newJob.NodeType)
			// TODO: Handle?
			continue
		}
		index := gatewayServer.roundRobinIndex[newJob.NodeType]
		stream := streams[index]

		// if round-robin index is out-of-bound: reset to 0
		if index+1 >= len(streams) {
			gatewayServer.roundRobinIndex[newJob.NodeType] = 0
		} else {
			gatewayServer.roundRobinIndex[newJob.NodeType]++
		}

		jobInput, err := json.Marshal(newJob.Input)
		if err != nil {
			log.Fatalf("failed to transform jobinput into json: %s", err.Error())
			continue
		}

		job := &pb.ActivatedJob{
			JobId:      newJob.ID,
			Type:       newJob.NodeType,
			WorkflowId: newJob.WorkflowID,
			Input:      string(jobInput),
		}
		(*stream).Send(&pb.ActivateJobResponse{
			Job: job,
		})
	}
}

func (gwServer *GatewayServer) keepAliveStream() {
	wait := make(chan any)
	<-wait
}

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
