del /Q bin
set GOOS=linux
go build -ldflags="-s -w" -o bin/lambda endpoints/aws/lambda.go