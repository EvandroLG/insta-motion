.SILENT:

build:
	go build -o bin/ ./...

test:
	go test ./...

clean:
	rm -rf bin/
