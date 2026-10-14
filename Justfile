[doc('Show available commands')]
@default:
    just --list

[doc('Run linter')]
@lint:
    echo "Lint..."
    command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint is not installed or in PATH"; exit 1; }
    GOARCH=wasm GOOS=wasip1 golangci-lint run
    # The P3 implementation is only compiled under this tag.
    GOARCH=wasm GOOS=wasip1 golangci-lint run --build-tags componentizego_async

[doc('Build both worlds')]
@build:
    echo "Build (wasip2)..."
    GOARCH=wasm GOOS=wasip1 go build ./...
    # The P3 implementation (*_p3.go) is only compiled under this tag, which
    # componentize-go sets for async worlds.
    echo "Build (wasip3, componentizego_async)..."
    GOARCH=wasm GOOS=wasip1 go build -tags componentizego_async ./...
