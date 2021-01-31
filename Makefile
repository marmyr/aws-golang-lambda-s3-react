.PHONY: build clean deploy

build:
	rm -f bin/*
	env GOOS=linux go build -ldflags="-s -w" -o bin/lambda endpoints/aws/lambda.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

