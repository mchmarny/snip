.PHONY: run mod build

# DEV
run:
	go run *.go -v

# BUILD
mod:
	go mod tidy
	go mod vendor

build: mod
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    	go build -a -tags netgo -ldflags '-w -extldflags "-static"' \
    	-mod vendor -o bin/snip-linux

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
    	go build -a -tags netgo -ldflags '-w -extldflags "-static"' \
    	-mod vendor -o bin/snip-mac

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
    	go build -a -tags netgo -ldflags '-w -extldflags "-static"' \
    	-mod vendor -o bin/snip-windows
