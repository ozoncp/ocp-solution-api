package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-solution-api/internal/api"
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-verdict-api"
)

const (
	grpcPort = ":7003"
)

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	// this Pings the database trying to connect
	db, err := sqlx.Connect("postgres", "user=postgres dbname=ozoncp sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	repo := repo.NewRepo(*db)
	const batchSize = 64
	desc.RegisterOcpVerdictApiServer(s, api.NewOcpVerdictApi(repo, batchSize))

	fmt.Printf("Verdict gRPC server is listening on localhost%v\n", grpcPort)
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
