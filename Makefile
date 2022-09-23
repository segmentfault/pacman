.PHONY: build clean test

GO_ENV=CGO_ENABLED=0
GO_FLAGS=-ldflags="-X main.BuildVersion=$(VERSION) -X 'main.BuildTime=`date`' -extldflags -static"
GO=$(GO_ENV) $(shell which go)

build:
	@$(GO) build ./...

test:
	@$(GO) test ./...

# clean all build result
clean:
	@$(GO) clean ./...
