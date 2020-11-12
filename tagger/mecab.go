// +build mecab

package tagger

import (
	"github.com/kaz/mecab-grpc/tagger/mecab"
)

func New(config Config) (Tagger, error) {
	return mecab.New(config)
}
