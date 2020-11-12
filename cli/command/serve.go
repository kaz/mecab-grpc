package command

import (
	"fmt"
	"net"

	"github.com/kaz/mecab-grpc/mecabpb"
	"github.com/kaz/mecab-grpc/remotetagger"
	"github.com/kaz/mecab-grpc/tagger"
	"google.golang.org/grpc"
)

type (
	Serve struct {
		Listen string            `short:"l" long:"listen"`
		Config map[string]string `short:"c" long:"config"`
	}
)

func (s *Serve) Execute(args []string) error {
	localTagger, err := tagger.New(s.Config)
	if err != nil {
		return fmt.Errorf("tagger.New: %w", err)
	}

	server := grpc.NewServer()
	mecabpb.RegisterMeCabServer(server, remotetagger.NewServer(localTagger))

	listener, err := net.Listen("tcp", s.Listen)
	if err != nil {
		return fmt.Errorf("new.Listen: %w", err)
	}

	fmt.Println("Listening on", listener.Addr())
	return server.Serve(listener)
}
