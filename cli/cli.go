package cli

import (
	"github.com/jessevdk/go-flags"
	"github.com/kaz/mecab-grpc/cli/command"
)

type (
	GlobalOptions struct{}
)

func Run() error {
	gOpts := &GlobalOptions{}
	parser := flags.NewParser(gOpts, flags.Default)

	parser.AddCommand("parse", "", "", &command.Parse{})
	parser.AddCommand("serve", "", "", &command.Serve{})

	if _, err := parser.Parse(); err != nil {
		if fe, ok := err.(*flags.Error); ok && fe.Type == flags.ErrHelp {
			return nil
		}
		return err
	}
	return nil
}
