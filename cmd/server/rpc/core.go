package rpc

import (
	"context"
	"github.com/matnich89/highscores-proto/core/v1"
	"log"
)

type CoreServiceServer struct {
	core.UnimplementedCoreServiceServer
}

func (w *CoreServiceServer) AddStats(_ context.Context, in *core.SendStatsRequest) (*core.SendStatsResponse, error) {
	log.Println(in.CharacterName)
	return &core.SendStatsResponse{
		Status: "OK",
	}, nil
}
