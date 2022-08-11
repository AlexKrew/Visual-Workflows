package gatewayserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	pb "workflows/gateway"

	"google.golang.org/grpc"
)

type GatewayServer struct {
	pb.UnimplementedGatewayServer
}

func (server *GatewayServer) CheckHealth(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	log.Printf("Received: %v", in.GetPing())
	return &pb.Pong{Pong: in.GetPing() + 1}, nil
}

func (server *GatewayServer) ActivateJob(ctx context.Context, in *pb.ActivateJobRequest) (*pb.ActivateJobResponse, error) {
	return nil, errors.New("not implemented")
}

func (server *GatewayServer) CompleteJob(ctx context.Context, in *pb.CompleteJobRequest) (*pb.CompleteJobResponse, error) {
	return nil, errors.New("not implemented")
}

func StartGatewayServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to start gateway server: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGatewayServer(server, &GatewayServer{})

	log.Printf("gateway server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failt to serve gateway server: %v", err)
	}
}
