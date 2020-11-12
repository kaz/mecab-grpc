package remotetagger

import (
	"context"
	"fmt"

	"github.com/kaz/mecab-grpc/mecabpb"
	"github.com/kaz/mecab-grpc/tagger"
)

type (
	TaggerServer struct {
		tagger tagger.Tagger

		mecabpb.UnimplementedMeCabServer
	}
)

func NewServer(t tagger.Tagger) *TaggerServer {
	return &TaggerServer{tagger: t}
}

func (s *TaggerServer) Parse(ctx context.Context, req *mecabpb.ParseRequest) (*mecabpb.ParseResponse, error) {
	parsed, err := s.tagger.Parse(req.GetInput())
	if err != nil {
		return nil, fmt.Errorf("tagger.Parse: %w", err)
	}
	return &mecabpb.ParseResponse{Output: parsed}, nil
}

func (s *TaggerServer) Close() error {
	return s.tagger.Close()
}
