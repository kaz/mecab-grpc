# mecab-grpc

[![PkgGoDev](https://pkg.go.dev/badge/github.com/kaz/mecab-grpc/mecabpb)](https://pkg.go.dev/github.com/kaz/mecab-grpc/mecabpb)

A server which provides Japanese tokenizing service over gRPC, powered with MeCab/NEologd.

## Usage

### [Server] Launch

```
$ docker built -t mecab-grpc .
....

$ docker run --rm -p 9000:9000 mecab-grpc
Listening on [::]:9000
```

You can edit `/usr/local/etc/mecabrc` to modify mecab's behavior.

Prebuilt image is also available on https://hub.docker.com/r/sekai/mecab-grpc

### [Client] Tokenize over gRPC

```
$ go build
....

$ echo "ハロー、世界！" | ./mecab-grpc parse --remote localhost:9000
ハロー  名詞,一般,*,*,*,*,ハロー,ハロー,ハロー
、      記号,読点,*,*,*,*,、,、,、
世界    名詞,一般,*,*,*,*,世界,セカイ,セカイ
！      記号,一般,*,*,*,*,！,！,！
EOS
```

### [Client] Embed in you app

This repository provides a Go API for communicating with mecab-grpc server.
See [pkg.go.dev](https://pkg.go.dev/github.com/kaz/mecab-grpc/mecabpb) or `/mecabpb` directory for detail.

You can also use `/mecabpb/mecab.proto` to generate gRPC codes for any other languages.

#### Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/kaz/mecab-grpc/mecabpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := mecabpb.NewMeCabClient(conn)
	resp, err := client.Parse(context.Background(), &mecabpb.ParseRequest{Input: "こんにちは、世界！"})
	if err != nil {
		panic(err)
	}

	fmt.Print(resp.Output)
}
```
