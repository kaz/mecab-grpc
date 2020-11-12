PB_DIR:=mecabpb

.PHONY: ${PB_DIR}
${PB_DIR}: ${PB_DIR}/mecab.pb.go ${PB_DIR}/mecab_grpc.pb.go

${PB_DIR}/mecab.pb.go: ${PB_DIR}/mecab.proto
	protoc --go_out=$(@D) $^

${PB_DIR}/mecab_grpc.pb.go: ${PB_DIR}/mecab.proto
	protoc --go-grpc_out=$(@D) $^
