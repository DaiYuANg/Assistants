all:

build:
	go build -o output/protocolassistant

clean:
	rm -rf output

dev:
	go run github.com/cosmtrek/air

fmt:
	go fmt ./...