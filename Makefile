BINARY_NAME=web/snake.wasm

all: build

build:
	GOARCH=wasm GOOS=js go build -o ${BINARY_NAME} cmd/wasm/*.go

clean:
	go clean
	rm ${BINARY_NAME}

test:
	GOARCH=wasm GOOS=js go test -v