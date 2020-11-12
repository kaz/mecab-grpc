package remotetagger

import (
	"context"
	"fmt"

	"github.com/kaz/mecab-grpc/pb"
	"github.com/kaz/mecab-grpc/tagger"
)

type (
	TaggerServer struct {
		tagger tagger.Tagger
	}
)

func NewServer(t tagger.Tagger) *TaggerServer {
	return &TaggerServer{tagger: t}
}

func (s *TaggerServer) Parse(ctx context.Context, req *pb.ParseRequest) (*pb.ParseResponse, error) {
	parsed, err := s.tagger.Parse(req.GetInput())
	if err != nil {
		return nil, fmt.Errorf("tagger.Parse: %w", err)
	}
	return &pb.ParseResponse{Output: parsed}, nil
}

func (s *TaggerServer) Close() error {
	return s.tagger.Close()
}
