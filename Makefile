BINARY_NAME=zvc
 
all: build

build:
	go build -o bin/${BINARY_NAME} main.go

clean:
	go clean
	rm bin/${BINARY_NAME}
