include .env
export

default: api run

run: api
	go run main.go

build: api bin/websrv

PROTO_SRC = $(wildcard api/*.proto)
PROTO_DST = $(patsubst api/%.proto, api/%.pb.go, $(PROTO_SRC))

api: $(PROTO_DST)

%.pb.go: %.proto
	protoc --go_out=paths=source_relative,plugins=grpc:$(dir $<) -I $(dir $<) -I api $<

clean:
	find api -name '*.pb.go' -delete
	find . -type f -and "(" -name main -or -name websrv -or -name simplify-godash ")" -delete

compile-js:
	npm --prefix js/ run build
