package main

import (
	"os"

	"github.com/kaz/mecab-grpc/cli"
)

func main() {
	if err := cli.Run(); err != nil {
		os.Exit(1)
	}
}
