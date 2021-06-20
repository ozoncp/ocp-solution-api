package main

import (
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	api "github.com/ozoncp/ocp-solution-api/internal/api"
	"github.com/ozoncp/ocp-solution-api/internal/repo"
	desc "github.com/ozoncp/ocp-solution-api/pkg/ocp-solution-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

const (
	grpcPort      = ":7002"
	prometeusPort = ":9001"
)

func runGrpc() error {
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
	desc.RegisterOcpSolutionApiServer(s, api.NewOcpSolutionApi(repo, batchSize))

	fmt.Printf("Solution gRPC server is listening on localhost%v\n", grpcPort)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	cfg := jaegercfg.Configuration{
		ServiceName: "ocp-solution-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	// Example logger and metrics factory. Use github.com/uber/jaeger-client-go/log
	// and github.com/uber/jaeger-lib/metrics respectively to bind to real logging and metrics
	// frameworks.
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	if err != nil {
		panic(err.Error())
	}

	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	go func() {
		// gRPC
		if err := runGrpc(); err != nil {
			log.Fatal(err)
		}
	}()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(prometeusPort, nil))
}
