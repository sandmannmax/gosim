run: build
	@bin/gosim

build:
	@go build -o bin/gosim .
