package command

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kaz/mecab-grpc/remotetagger"
	"github.com/kaz/mecab-grpc/tagger"
)

type (
	Parse struct {
		Local  bool   `short:"l" long:"local"`
		Remote string `short:"r" long:"remote"`

		Config map[string]string `short:"c" long:"config"`
	}
)

func (p *Parse) Execute(args []string) error {
	var t tagger.Tagger
	var err error

	if p.Local {
		t, err = tagger.New(p.Config)
		if err != nil {
			return fmt.Errorf("tagger.New: %w", err)
		}
	} else if p.Remote != "" {
		t, err = remotetagger.New(p.Remote)
		if err != nil {
			return fmt.Errorf("remotetagger.New: %w", err)
		}
	} else {
		return fmt.Errorf("--local or --remote must be specified")
	}
	defer t.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parsed, err := t.Parse(scanner.Text())
		if err != nil {
			return fmt.Errorf("Parse: %w", err)
		}

		fmt.Print(parsed)
	}

	return nil
}
