package tagger

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)

type (
	LocalTagger struct {
		tagger mecab.MeCab
	}
)

func New(config map[string]string) (*LocalTagger, error) {
	tagger, err := mecab.New(config)
	if err != nil {
		return nil, fmt.Errorf("mecab.New: %w", err)
	}
	return &LocalTagger{tagger: tagger}, nil
}

func (t *LocalTagger) Parse(input string) (string, error) {
	parsed, err := t.tagger.Parse(input)
	if err != nil {
		return "", fmt.Errorf("tagger.Parse: %w", err)
	}
	return parsed, nil
}

func (t *LocalTagger) Close() error {
	t.tagger.Destroy()
	return nil
}
