package main

import (
	"highscores-core/cmd/server/rpc"
	"highscores-core/internal/logger"
)

func main() {
	cfg := NewConfig(":50051")
	app := &application{
		logger: &logger.Logger{},
		config: cfg,
		server: &rpc.CoreServiceServer{},
	}
	app.listen()
}
