// +build !mecab

package tagger

import "github.com/kaz/mecab-grpc/tagger/mock"

func New(config Config) (Tagger, error) {
	return mock.New(), nil
}
