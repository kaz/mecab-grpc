package mecab

import (
	"fmt"
)

type (
	Config = map[string]string

	MecabTagger struct {
		pool *MecabPool
	}
)

func New(config Config) (*MecabTagger, error) {
	return &MecabTagger{pool: NewPool(config)}, nil
}

func (t *MecabTagger) Parse(input string) (string, error) {
	tagger, err := t.pool.Get()
	if err != nil {
		return "", fmt.Errorf("pool.Get: %w", err)
	}

	parsed, err := tagger.Parse(input)
	if err != nil {
		return "", fmt.Errorf("tagger.Parse: %w", err)
	}

	t.pool.Put(tagger)
	return parsed, nil
}

func (t *MecabTagger) Close() error {
	return nil
}
