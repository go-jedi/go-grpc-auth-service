package httpserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-jedi/auth/config"
	protoservice "github.com/go-jedi/auth/gen/proto/service/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultHost = "127.0.0.1"
	defaultPort = 50052
)

type HTTPServer struct {
	host string
	port int
}

func (hs *HTTPServer) init() error {
	if hs.host == "" {
		hs.host = defaultHost
	}

	if hs.port == 0 {
		hs.port = defaultPort
	}

	return nil
}

func NewHTTPServer(cfg config.HTTPServerConfig) (*HTTPServer, error) {
	hs := &HTTPServer{
		host: cfg.Host,
		port: cfg.Port,
	}

	if err := hs.init(); err != nil {
		return nil, err
	}

	return hs, nil
}

func (hs *HTTPServer) Start(ctx context.Context, grpcPort int) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := protoservice.RegisterAuthV1HandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts); err != nil {
		return err
	}

	if err := protoservice.RegisterUserV1HandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", grpcPort), opts); err != nil {
		return err
	}

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", hs.port),
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("http server listening at: %d", hs.port)

	return s.ListenAndServe()
}
