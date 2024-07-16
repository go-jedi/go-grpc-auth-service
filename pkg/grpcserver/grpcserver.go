package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/go-jedi/auth/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 50053
)

type GRPCServer struct {
	Server *grpc.Server

	host string
	port int
}

func (gs *GRPCServer) init() error {
	if gs.host == "" {
		gs.host = defaultHost
	}

	if gs.port == 0 {
		gs.port = defaultPort
	}

	return nil
}

func NewGRPCServer(cfg config.GRPCServerConfig) (*GRPCServer, error) {
	gs := &GRPCServer{
		host: cfg.Host,
		port: cfg.Port,
	}

	if err := gs.init(); err != nil {
		return nil, err
	}

	gs.Server = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	reflection.Register(gs.Server)

	return gs, nil
}

// Start grpc server.
func (gs *GRPCServer) Start() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", gs.port))
	if err != nil {
		return err
	}

	log.Printf("gRPC server listening at: %d", gs.port)

	return gs.Server.Serve(l)
}
