# generate version number
version=$(shell git describe --tags --long --always --dirty|sed 's/^v//')
binfile=mqpp

all: vendor | glide.lock
	go build -ldflags "-X main.version=$(version)"
	-@go fmt

static: glide.lock vendor
	go build -ldflags "-X main.version=$(version) -extldflags \"-static\"" -o $(binfile).static $(binfile).go

arm:
	GOARCH=arm go build  -ldflags "-X main.version=$(version) -extldflags \"-static\"" -o $(binfile).arm $(binfile).go
	GOARCH=arm64 go build  -ldflags "-X main.version=$(version) -extldflags \"-static\"" -o $(binfile).arm64 $(binfile).go
clean:
	rm -rf vendor
	rm -rf _vendor
vendor: glide.lock
	glide install && touch vendor
glide.lock: glide.yaml
	glide update && touch glide.lock
glide.yaml:
version:
	@echo $(version)
