.PHONY: build clean deploy

build:
	rm -f bin/*
	env GOOS=linux go build -ldflags="-s -w" -o bin/aws endpoints/aws/aws.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

