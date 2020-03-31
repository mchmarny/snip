RELEASE_VERSION=v0.4.4

.PHONY: run mod build

# DEV

test: mod
	go test ./... -v

install: mod
	go install

list:
	go list -m all

update:
	go get -u ./...

# BUILD
mod:
	go mod tidy
	go mod vendor

build: mod
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags netgo -ldflags \
	'-w -extldflags "-static" -X main.AppVersion=${RELEASE_VERSION}' \
	-mod vendor -o bin/snip-mac-${RELEASE_VERSION}

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags \
	'-w -extldflags "-static" -X main.AppVersion=${RELEASE_VERSION}' \
	-mod vendor -o bin/snip-linux-${RELEASE_VERSION}

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -tags netgo -ldflags \
	'-w -extldflags "-static" -X main.AppVersion=${RELEASE_VERSION}' \
	-mod vendor -o bin/snip-windows-${RELEASE_VERSION}

tag:
	git tag "release-${RELEASE_VERSION}"
	git push origin "release-${RELEASE_VERSION}"

assets:
	script/release-asset.sh "${RELEASE_VERSION}"

