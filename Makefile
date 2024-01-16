GOFILES := $(shell find ./src/go -name \*.go -print)
GOCLIENTFILES := $(shell find ./src/go-client -name \*.go -print)
GOOS=darwin
GOARCH=arm64

build/go-server.$(GOOS).$(GOARCH): $(GOFILES) src/go/pb/*.pb.go src/go/pb/*_grpc.pb.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/go-server.$(GOOS).$(GOARCH) ./src/go

src/go/pb/%.pb.go src/go/pb/%_grpc.pb.go: src/protos/%.proto
	protoc --go_out=src/go/pb --go_opt=paths=source_relative \
		--go-grpc_out=src/go/pb --go-grpc_opt=paths=source_relative \
		--proto_path src/protos/ \
		src/protos/*.proto

build/go-client.$(GOOS).$(GOARCH): $(GOCLIENTFILES) src/go-client/pb/*.pb.go src/go-client/pb/*_grpc.pb.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o build/go-client.$(GOOS).$(GOARCH) ./src/go-client

src/go-client/pb/%.pb.go src/go-client/pb/%_grpc.pb.go: src/protos/%.proto
	protoc --go_out=src/go-client/pb --go_opt=paths=source_relative \
		--go-grpc_out=src/go-client/pb --go-grpc_opt=paths=source_relative \
		--proto_path src/protos/ \
		src/protos/*.proto
