package main

import (
	"github.com/matnich89/highscores-proto/core/v1"
	"google.golang.org/grpc"
	"highscores-core/cmd/server/rpc"
	"highscores-core/internal/logger"
	"log"
	"net"
)

type application struct {
	logger *logger.Logger
	config *Config
	server *rpc.CoreServiceServer
}

func (a *application) listen() {
	lis, err := net.Listen("tcp", a.config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	core.RegisterCoreServiceServer(s, a.server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
