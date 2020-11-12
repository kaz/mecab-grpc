# mecab-grpc

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
See `/mecabpb` directory for detail.

You can also use `/mecabpb/mecab.proto` to generate gRPC codes for any other languages.
