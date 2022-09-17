PACKAGES=./packages...
BINARY_NAME=automata-encoder
BINARY_DEST=./bin/${BINARY_NAME}

.PHONY: test
test: 
	go test ${PACKAGES}

.PHONY: build
build: build-linux build-windows build-darwin

.PHONY: build-linux
build-linux: 
	GOOS=linux go build -o ${BINARY_DEST}-linux

.PHONY: build-windows
build-windows: 
	GOOS=windows go build -o ${BINARY_DEST}.exe

.PHONY: build-darwin
build-darwin: 
	GOOS=darwin go build -o ${BINARY_DEST}-darwin