RELEASE_VERSION=v0.3.1

.PHONY: run mod build

# DEV

test: mod
	go test ./... -v

install: mod
	go install

# BUILD
mod:
	go mod tidy
	go mod vendor

build: mod
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
    	go build -a -tags netgo -ldflags \
		'-w -extldflags "-static" -X main.AppVersion=${RELEASE_VERSION}' \
    	-mod vendor -o bin/snip-mac

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    	go build -a -tags netgo -ldflags \
		'-w -extldflags "-static" -X main.AppVersion=${RELEASE_VERSION}' \
    	-mod vendor -o bin/snip-linux

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
    	go build -a -tags netgo -ldflags \
		'-w -extldflags "-static" -X main.AppVersion=${RELEASE_VERSION}' \
    	-mod vendor -o bin/snip-windows

tag:
	git tag "release-${RELEASE_VERSION}"
	git push origin "release-${RELEASE_VERSION}"
	git log --oneline