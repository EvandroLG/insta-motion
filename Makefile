.PHONY: run build test clean
.SILENT:

run:
	go run ./...

build:
	mkdir -p bin
	go build -o bin/ ./...

test:
	go test ./...

clean:
	rm -rf bin/
