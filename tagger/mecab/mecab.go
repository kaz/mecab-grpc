package mecab

import (
	"fmt"

	"github.com/shogo82148/go-mecab"
)

type (
	MecabTagger struct {
		tagger mecab.MeCab
	}
)

func New(config map[string]string) (*MecabTagger, error) {
	tagger, err := mecab.New(config)
	if err != nil {
		return nil, fmt.Errorf("mecab.New: %w", err)
	}
	return &MecabTagger{tagger: tagger}, nil
}

func (t *MecabTagger) Parse(input string) (string, error) {
	parsed, err := t.tagger.Parse(input)
	if err != nil {
		return "", fmt.Errorf("tagger.Parse: %w", err)
	}
	return parsed, nil
}

func (t *MecabTagger) Close() error {
	t.tagger.Destroy()
	return nil
}
