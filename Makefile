default: out/example

clean:
	rm -rf out

test: *.go
	go test ./...

out/example: lab2/implementation.go lab2/handler.go cmd/example/main.go
	mkdir -p out
	go build -o out/example ./cmd/example
