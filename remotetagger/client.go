package remotetagger

import (
	"context"
	"fmt"

	"github.com/kaz/mecab-grpc/pb"
	"google.golang.org/grpc"
)

type (
	RemoteTagger struct {
		conn   *grpc.ClientConn
		client pb.MeCabClient
	}
)

func New(target string, opts ...grpc.DialOption) (*RemoteTagger, error) {
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, fmt.Errorf("grpc.Dial: %w", err)
	}
	return &RemoteTagger{
		conn:   conn,
		client: pb.NewMeCabClient(conn),
	}, nil
}

func (t *RemoteTagger) Parse(input string) (string, error) {
	return t.ParseWithContext(context.Background(), input)
}
func (t *RemoteTagger) ParseWithContext(ctx context.Context, input string) (string, error) {
	resp, err := t.client.Parse(ctx, &pb.ParseRequest{Input: input})
	if err != nil {
		return "", fmt.Errorf("client.Parse: %w", err)
	}
	return resp.GetOutput(), nil
}

func (t *RemoteTagger) Close() error {
	return t.conn.Close()
}
