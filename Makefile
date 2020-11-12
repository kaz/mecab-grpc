.PHONY: pb
pb: pb/mecab.pb.go pb/mecab_grpc.pb.go

pb/mecab.pb.go: pb/mecab.proto
	protoc --go_out=$(@D) $^

pb/mecab_grpc.pb.go: pb/mecab.proto
	protoc --go-grpc_out=$(@D) $^
