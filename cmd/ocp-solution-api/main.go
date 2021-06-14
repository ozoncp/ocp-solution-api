package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"fmt"
	"log"
	"net"

	api "github.com/ozoncp/ocp-solution-api/internal/api"
	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-solution-api"
)

const (
	grpcPort = ":7002"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterOcpSolutionApiServer(s, api.NewOcpSolutionApi())

	fmt.Printf("server is listening on localhost%v\n", grpcPort)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
