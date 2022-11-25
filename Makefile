.PHONY: build

clean:
	@go clean
	@rm -rf dist/

build:
	@goreleaser build --single-target --snapshot --rm-dist

all: clean build

release:
	@echo "Don't forget to set GITHUB_TOKEN environment variable"
	@goreleaser release --rm-dist