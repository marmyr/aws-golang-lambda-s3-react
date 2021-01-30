.PHONY: build clean deploy

build:
	rm -f bin/*
	env GOOS=linux go build -ldflags="-s -w" -o bin/hello endpoints/aws/hello.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

